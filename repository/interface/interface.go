package _interface

import (
	"context"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
)

type ReadWrite interface {
	WriteTag(ctx context.Context, req *pb.Tag) (*pb.Tag, error)
	ModifyTag(ctx context.Context, req *pb.Tag) (*pb.Tag, error)
	RemoveTag(ctx context.Context, req *pb.Select) error
	ReadTags(ctx context.Context) (*pb.Tags, error)

	WriteTopic(ctx context.Context, req *pb.Topic) (*pb.Topic, error)
	ModifyTopic(ctx context.Context, req *pb.Topic) (*pb.Topic, error)
	RemoveTopic(ctx context.Context, req *pb.Select) error
	ReadTopics(ctx context.Context) (*pb.Topics, error)

	WriteNews(ctx context.Context, req *pb.News) (*pb.News, error)
	ModifyNews(ctx context.Context, req *pb.News) (*pb.News, error)
	RemoveNews(ctx context.Context, req *pb.Select) error
	ReadNewses(ctx context.Context) (*pb.Newses, error)
}

type Cache interface {
	SetTag(ctx context.Context, tag *pb.Tag) error
	UnsetTag(ctx context.Context, id string) error
	GetTags(ctx context.Context) (*pb.Tags, error)
	ReloadTags(ctx context.Context, tags *pb.Tags) error

	SetTopic(ctx context.Context, tag *pb.Topic) error
	UnsetTopic(ctx context.Context, id string) error
	GetTopics(ctx context.Context) (*pb.Topics, error)
	ReloadTopics(ctx context.Context, topics *pb.Topics) error

	SetNews(ctx context.Context, tag *pb.News) error
	UnsetNews(ctx context.Context, id string) error
	GetNewses(ctx context.Context) (*pb.News, error)
	ReloadNewses(ctx context.Context, newses *pb.Newses) error
}
