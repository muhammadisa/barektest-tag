package main

import (
	"context"
	oczipkin "contrib.go.opencensus.io/exporter/zipkin"
	"fmt"
	"github.com/go-kit/kit/log/level"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/muhammadisa/barektest-tag/constant"
	ep "github.com/muhammadisa/barektest-tag/endpoint"
	"github.com/muhammadisa/barektest-tag/gvars"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
	"github.com/muhammadisa/barektest-tag/repository"
	"github.com/muhammadisa/barektest-tag/service"
	"github.com/muhammadisa/barektest-tag/transport"
	"github.com/muhammadisa/barektest-util/cb"
	"github.com/muhammadisa/barektest-util/dbc"
	"github.com/muhammadisa/barektest-util/hdr"
	"github.com/muhammadisa/barektest-util/lgr"
	"github.com/muhammadisa/barektest-util/vlt"
	"github.com/openzipkin/zipkin-go"
	httpreporter "github.com/openzipkin/zipkin-go/reporter/http"
	"github.com/soheilhy/cmux"
	"go.opencensus.io/trace"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

// ServeGRPC serving GRPC server and will be merged using MergeServer function
func ServeGRPC(listener net.Listener, service pb.TagServiceServer, serverOptions []grpc.ServerOption) error {
	level.Info(gvars.Log).Log(lgr.LogInfo, "initialize grpc server")

	var grpcServer *grpc.Server
	if len(serverOptions) > 0 {
		grpcServer = grpc.NewServer(serverOptions...)
	} else {
		grpcServer = grpc.NewServer()
	}
	pb.RegisterTagServiceServer(grpcServer, service)
	return grpcServer.Serve(listener)
}

// ServeHTTP serving HTTP server and will be merged using MergeServer function
func ServeHTTP(listener net.Listener, service pb.TagServiceServer) error {
	level.Info(gvars.Log).Log(lgr.LogInfo, "initialize rest server")

	mux := runtime.NewServeMux()
	err := pb.RegisterTagServiceHandlerServer(context.Background(), mux, service)
	if err != nil {
		return err
	}
	srv := &http.Server{Handler: hdr.CORS(mux)}
	return srv.Serve(listener)
}

// MergeServer start ServeGRPC and ServeHTTP concurrently
func MergeServer(service pb.TagServiceServer, serverOptions []grpc.ServerOption) {
	port := fmt.Sprintf(":%s", "8010")
	listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(listener)
	grpcListener := m.MatchWithWriters(cmux.HTTP2MatchHeaderFieldSendSettings(
		"content-type", "application/grpc",
	))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return ServeGRPC(grpcListener, service, serverOptions) })
	g.Go(func() error { return ServeHTTP(httpListener, service) })
	g.Go(func() error { return m.Serve() })

	log.Fatal(g.Wait())
}

func main() {
	gvars.Log = lgr.Create(constant.ServiceName)

	level.Info(gvars.Log).Log(lgr.LogInfo, "service started")

	ctx := context.Background()
	defer ctx.Done()

	vault, err := vlt.NewVLT("myroot", "http://localhost:8300", "/secret")
	if err != nil {
		panic(err)
	}

	reporter := httpreporter.NewReporter(vault.Get("/tracer_conf:url"))
	localEndpoint, _ := zipkin.NewEndpoint(constant.ServiceName, ":0")
	exporter := oczipkin.NewExporter(reporter, localEndpoint)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
	trace.RegisterExporter(exporter)
	trcr := trace.DefaultTracer

	var repoConf repository.RepoConf
	{
		repoConf.SQL = dbc.Config{
			Username: vault.Get("/sql_database:username"),
			Password: vault.Get("/sql_database:password"),
			Host:     vault.Get("/sql_database:host"),
			Port:     vault.Get("/sql_database:port"),
			Name:     vault.Get("/sql_database:db"),
		}
		repoConf.Cache = dbc.Config{
			Password: "95798588",
			Host:     "localhost",
			Port:     "6379",
		}
	}

	tagRepo, err := repository.NewRepository(ctx, repoConf, trcr)
	if err != nil {
		panic(err)
	}

	err = cb.StartHystrix(10, constant.ServiceName)
	if err != nil {
		panic(err)
	}

	tagSvc := service.NewUsecases(*tagRepo, trcr, vault)

	tagEp, err := ep.NewTagEndpoint(tagSvc, gvars.Log, vault)
	if err != nil {
		panic(err)
	}

	server := transport.NewTagServer(tagEp)
	MergeServer(server, nil)
}
