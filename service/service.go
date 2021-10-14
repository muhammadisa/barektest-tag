package service

import (
	"context"
	"fmt"
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

func (s service) AddTag(ctx context.Context, tag *pb.Tag) (res *emptypb.Empty, err error) {
	newTag, err := s.repo.ReadWriter.WriteTag(ctx, tag)
	if err != nil {
		return nil, err
	}
	return nil, s.repo.CacheReadWriter.SetTag(ctx, newTag)
}

func (s service) EditTag(ctx context.Context, tag *pb.Tag) (res *emptypb.Empty, err error) {
	err = s.repo.CacheReadWriter.UnsetTag(ctx, tag.Id)
	if err != nil {
		return nil, err
	}
	updatedTag, err := s.repo.ReadWriter.ModifyTag(ctx, tag)
	if err != nil {
		return nil, err
	}
	return nil, s.repo.CacheReadWriter.SetTag(ctx, updatedTag)
}

func (s service) DeleteTag(ctx context.Context, selectTag *pb.SelectTag) (res *emptypb.Empty, err error) {
	err = s.repo.ReadWriter.RemoveTag(ctx, selectTag)
	if err != nil {
		return nil, err
	}
	err = s.repo.CacheReadWriter.UnsetTag(ctx, selectTag.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s service) GetTags(ctx context.Context, _ *emptypb.Empty) (res *pb.Tags, err error) {
	res, err = s.repo.CacheReadWriter.GetTags(ctx)
	if err != nil {
		return nil, err
	}
	if len(res.Tags) == 0 {
		fmt.Println("from database")
		res, err = s.repo.ReadWriter.ReadTags(ctx)
		if err != nil {
			return nil, err
		}
		err = s.repo.CacheReadWriter.ReloadTags(ctx, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	}
	fmt.Println("from cache")
	return res, nil
}

func NewUsecases(repo _repointerface.Repository, tracer trace.Tracer, vault vlt.VLT) _interface.Service {
	return &service{
		tracer: tracer,
		repo:   repo,
		vault:  vault,
	}
}
