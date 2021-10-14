package _interface

import (
	"context"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
)

type ReadWrite interface {
	WriteTag(ctx context.Context, req *pb.Tag) error
	ModifyTag(ctx context.Context, req *pb.Tag) error
	RemoveTag(ctx context.Context, req *pb.SelectTag) error
	ReadTags(ctx context.Context) (*pb.Tags, error)
}
