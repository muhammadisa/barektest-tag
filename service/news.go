package service

import (
	"context"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s service) AddNews(ctx context.Context, news *pb.News) (res *emptypb.Empty, err error) {
	newNews, err := s.repo.ReadWriter.WriteNews(ctx, news)
	if err != nil {
		return nil, err
	}

	err = s.repo.ReadWriter.WriteNewsTags(ctx, newNews.Id, news.NewsTagIds)
	if err != nil {
		return nil, err
	}

	return nil, s.repo.CacheReadWriter.SetNews(ctx, newNews)
}

func (s service) EditNews(ctx context.Context, news *pb.News) (res *emptypb.Empty, err error) {
	err = s.repo.CacheReadWriter.UnsetNews(ctx, news.Id)
	if err != nil {
		return nil, err
	}
	updatedNews, err := s.repo.ReadWriter.ModifyNews(ctx, news)
	if err != nil {
		return nil, err
	}
	return nil, s.repo.CacheReadWriter.SetNews(ctx, updatedNews)
}

func (s service) DeleteNews(ctx context.Context, selectNews *pb.Select) (res *emptypb.Empty, err error) {
	err = s.repo.ReadWriter.RemoveNews(ctx, selectNews)
	if err != nil {
		return nil, err
	}
	err = s.repo.CacheReadWriter.UnsetNews(ctx, selectNews.Id)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (s service) GetNewses(ctx context.Context, filters *pb.Filters) (res *pb.Newses, err error) {
	return s.repo.ReadWriter.ReadNewses(ctx)
}
