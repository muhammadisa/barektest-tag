package service

import (
	"context"
	"fmt"
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
		_ = s.repo.ReadWriter.RemoveNews(ctx, &pb.Select{Id: newNews.Id})
		return nil, err
	}
	return nil, s.repo.CacheReadWriter.InvalidateNewses(ctx)
}

func (s service) EditNews(ctx context.Context, news *pb.News) (res *emptypb.Empty, err error) {
	_ = s.repo.CacheReadWriter.UnsetNews(ctx, news.Id)
	oldNews, err := s.repo.ReadWriter.ModifyNews(ctx, news)
	if err != nil {
		return nil, err
	}
	err = s.repo.ReadWriter.WriteNewsTags(ctx, oldNews.Id, news.NewsTagIds)
	if err != nil {
		return nil, err
	}
	return nil, s.repo.CacheReadWriter.InvalidateNewses(ctx)
}

func (s service) DeleteNews(ctx context.Context, selectNews *pb.Select) (res *emptypb.Empty, err error) {
	_ = s.repo.CacheReadWriter.UnsetNews(ctx, selectNews.Id)
	err = s.repo.ReadWriter.RemoveNews(ctx, selectNews)
	if err != nil {
		return nil, err
	}
	return nil, s.repo.CacheReadWriter.InvalidateNewses(ctx)
}

func (s service) noneFilter(ctx context.Context) (res *pb.Newses, err error) {
	if reload := s.repo.CacheReadWriter.ReloadRequired(ctx); reload {
		fmt.Println("none - reload")
		res, err = s.repo.ReadWriter.ReadNewses(ctx)
		if err != nil {
			return nil, err
		}
		err = s.repo.CacheReadWriter.ReloadNewses(ctx, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		fmt.Println("none - no reload")
		return s.repo.CacheReadWriter.GetNewses(ctx)
	}
}

func (s service) statusFilter(ctx context.Context, status int32) (res *pb.Newses, err error) {
	if reload := s.repo.CacheReadWriter.ReloadRequired(ctx); reload {
		fmt.Println("status - reload")
		res, err = s.repo.ReadWriter.ReadNewsesByStatus(ctx, status)
		if err != nil {
			return nil, err
		}
		err = s.repo.CacheReadWriter.ReloadNewses(ctx, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		fmt.Println("status - no reload")
		return s.repo.CacheReadWriter.GetNewses(ctx)
	}
}

func (s service) topicIdFilter(ctx context.Context, topicID string) (res *pb.Newses, err error) {
	if reload := s.repo.CacheReadWriter.ReloadRequired(ctx); reload {
		fmt.Println("topic_id - reload")
		res, err = s.repo.ReadWriter.ReadNewsesByTopicID(ctx, topicID)
		if err != nil {
			return nil, err
		}
		err = s.repo.CacheReadWriter.ReloadNewses(ctx, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		fmt.Println("topic_id - no reload")
		return s.repo.CacheReadWriter.GetNewses(ctx)
	}
}

func (s service) statusAndTopicIdFilter(ctx context.Context, filters *pb.Filters) (res *pb.Newses, err error) {
	if reload := s.repo.CacheReadWriter.ReloadRequired(ctx); reload {
		fmt.Println("status topic_id - reload")
		res, err = s.repo.ReadWriter.ReadNewsesByStatusAndTopicID(ctx, filters.Status, filters.TopicId)
		if err != nil {
			return nil, err
		}
		err = s.repo.CacheReadWriter.ReloadNewses(ctx, res)
		if err != nil {
			return nil, err
		}
		return res, nil
	} else {
		fmt.Println("status topic_id - no reload")
		return s.repo.CacheReadWriter.GetNewses(ctx)
	}
}

func (s service) GetNewses(ctx context.Context, filters *pb.Filters) (res *pb.Newses, err error) {
	if filters.TopicId != "" && filters.Status != 0 {
		return s.statusAndTopicIdFilter(ctx, filters)
	}
	if filters.Status != 0 {
		return s.statusFilter(ctx, filters.Status)
	}
	if filters.TopicId != "" {
		return s.topicIdFilter(ctx, filters.TopicId)
	}
	return s.noneFilter(ctx)
}
