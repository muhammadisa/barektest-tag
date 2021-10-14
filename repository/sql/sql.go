package sql

import (
	"context"
	"database/sql"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
	_interface "github.com/muhammadisa/barektest-tag/repository/interface"
	"github.com/muhammadisa/barektest-util/dbc"
	"go.opencensus.io/trace"
	"sync"
)

var mutex = &sync.RWMutex{}

type readWrite struct {
	tracer trace.Tracer
	db     *sql.DB
}

func (r *readWrite) WriteTag(ctx context.Context, req *pb.Tag) error {
	panic("implement me")
}

func (r *readWrite) ModifyTag(ctx context.Context, req *pb.Tag) error {
	panic("implement me")
}

func (r *readWrite) RemoveTag(ctx context.Context, req *pb.SelectTag) error {
	panic("implement me")
}

func (r *readWrite) ReadTags(ctx context.Context) (*pb.Tags, error) {
	panic("implement me")
}

func NewSQL(config dbc.Config, tracer trace.Tracer) (_interface.ReadWrite, error) {
	sqlDB, err := dbc.OpenDB(config)
	if err != nil {
		return nil, err
	}
	return &readWrite{
		db:     sqlDB,
		tracer: tracer,
	}, nil
}

