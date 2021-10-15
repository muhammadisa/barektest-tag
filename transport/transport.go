package transport

import (
	"context"
	kitoc "github.com/go-kit/kit/tracing/opencensus"
	grpctransport "github.com/go-kit/kit/transport/grpc"
	ep "github.com/muhammadisa/barektest-tag/endpoint"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

type grpcTagServer struct {
	addTag    grpctransport.Handler
	editTag   grpctransport.Handler
	deleteTag grpctransport.Handler
	getTags   grpctransport.Handler

	addTopic    grpctransport.Handler
	editTopic   grpctransport.Handler
	deleteTopic grpctransport.Handler
	getTopics   grpctransport.Handler
}

func (g grpcTagServer) AddTopic(ctx context.Context, req *pb.Topic) (*emptypb.Empty, error) {
	_, res, err := g.addTopic.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcTagServer) EditTopic(ctx context.Context, req *pb.Topic) (*emptypb.Empty, error) {
	_, res, err := g.editTopic.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcTagServer) DeleteTopic(ctx context.Context, req *pb.Select) (*emptypb.Empty, error) {
	_, res, err := g.deleteTopic.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcTagServer) GetTopics(ctx context.Context, req *emptypb.Empty) (*pb.Topics, error) {
	_, res, err := g.getTopics.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Topics), nil
}

// ..

func (g grpcTagServer) AddTag(ctx context.Context, req *pb.Tag) (*emptypb.Empty, error) {
	_, res, err := g.addTag.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcTagServer) EditTag(ctx context.Context, req *pb.Tag) (*emptypb.Empty, error) {
	_, res, err := g.editTag.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcTagServer) DeleteTag(ctx context.Context, req *pb.Select) (*emptypb.Empty, error) {
	_, res, err := g.deleteTag.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*emptypb.Empty), nil
}

func (g grpcTagServer) GetTags(ctx context.Context, req *emptypb.Empty) (*pb.Tags, error) {
	_, res, err := g.getTags.ServeGRPC(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.(*pb.Tags), nil
}

func NewTagServer(endpoints ep.TagEndpoint) pb.TagServiceServer {
	options := []grpctransport.ServerOption{
		kitoc.GRPCServerTrace(),
	}
	return &grpcTagServer{
		addTag: grpctransport.NewServer(
			endpoints.AddTagEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		editTag: grpctransport.NewServer(
			endpoints.EditTagEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		deleteTag: grpctransport.NewServer(
			endpoints.DeleteTagEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		getTags: grpctransport.NewServer(
			endpoints.GetTagsEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		//..
		addTopic: grpctransport.NewServer(
			endpoints.AddTopicEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		editTopic: grpctransport.NewServer(
			endpoints.EditTopicEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		deleteTopic: grpctransport.NewServer(
			endpoints.DeleteTopicEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
		getTopics: grpctransport.NewServer(
			endpoints.GetTopicsEndpoint,
			decodeRequest,
			encodeResponse,
			options...,
		),
	}
}

func decodeRequest(_ context.Context, request interface{}) (interface{}, error) {
	return request, nil
}

func encodeResponse(_ context.Context, response interface{}) (interface{}, error) {
	return response, nil
}
