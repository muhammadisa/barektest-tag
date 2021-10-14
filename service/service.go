package service

import (
	"context"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
	_repointerface "github.com/muhammadisa/barektest-tag/repository"
	_interface "github.com/muhammadisa/barektest-tag/service/interface"
	"github.com/muhammadisa/barektest-util/vlt"
	"go.opencensus.io/trace"
	"google.golang.org/protobuf/types/known/emptypb"
)

type service struct {
	vault  vlt.VLT
	tracer trace.Tracer
	repo   _repointerface.Repository
}

func (s service) AddTag(ctx context.Context, tag *pb.Tag) (*emptypb.Empty, error) {
	return nil, s.repo.ReadWriter.WriteTag(ctx, tag)
}

func (s service) EditTag(ctx context.Context, tag *pb.Tag) (*emptypb.Empty, error) {
	return nil, s.repo.ReadWriter.ModifyTag(ctx, tag)
}

func (s service) DeleteTag(ctx context.Context, selectTag *pb.SelectTag) (*emptypb.Empty, error) {
	return nil, s.repo.ReadWriter.RemoveTag(ctx, selectTag)
}

func (s service) GetTags(ctx context.Context, _ *emptypb.Empty) (*pb.Tags, error) {
	return s.repo.ReadWriter.ReadTags(ctx)
}

func NewUsecases(repo _repointerface.Repository, tracer trace.Tracer, vault vlt.VLT) _interface.Service {
	return &service{
		tracer: tracer,
		repo:   repo,
		vault:  vault,
	}
}
