package _interface

import (
	"context"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
)

type ReadWrite interface {
	WriteTag(ctx context.Context, req *pb.Tag) (*pb.Tag, error)
	ModifyTag(ctx context.Context, req *pb.Tag) (*pb.Tag, error)
	RemoveTag(ctx context.Context, req *pb.SelectTag) error
	ReadTags(ctx context.Context) (*pb.Tags, error)
}

/**
map[string]interface
{
	"tag_uuid": tag_detail
}
*/

type Cache interface {
	SetTag(ctx context.Context, tag *pb.Tag) error
	UnsetTag(ctx context.Context, id string) error
	GetTags(ctx context.Context) (*pb.Tags, error)
	ReloadTags(ctx context.Context, tags *pb.Tags) error
}
