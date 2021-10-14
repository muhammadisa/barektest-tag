package cache

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
	_interface "github.com/muhammadisa/barektest-tag/repository/interface"
	"github.com/muhammadisa/barektest-util/dbc"
	"go.opencensus.io/trace"
)

type cache struct {
	tracer trace.Tracer
	redis  redis.Cmdable
}

func (c *cache) ReloadTags(ctx context.Context, tags *pb.Tags) (err error) {
	data := make(map[string]interface{})
	for _, tag := range tags.Tags {
		tagByte, err := json.Marshal(tag)
		if err != nil {
			return err
		}
		data[tag.Id] = string(tagByte)
	}
	c.redis.HMSet(ctx, "tags", data)
	return nil
}

func (c *cache) UnsetTag(ctx context.Context, id string) (err error) {
	return c.redis.HDel(ctx, "tags", id).Err()
}

func (c *cache) SetTag(ctx context.Context, tag *pb.Tag) (err error) {
	tagByte, err := json.Marshal(tag)
	if err != nil {
		return err
	}
	return c.redis.HSetNX(ctx, "tags", tag.Id, string(tagByte)).Err()
}

func (c *cache) GetTags(ctx context.Context) (res *pb.Tags, err error) {
	var tags pb.Tags
	tagsMap := c.redis.HGetAll(ctx, "tags").Val()
	for _, v := range tagsMap {
		var tag pb.Tag
		err = json.Unmarshal([]byte(v), &tag)
		if err != nil {
			return res, err
		}
		tags.Tags = append(tags.Tags, &tag)
	}
	return &tags, nil
}

func NewCache(config dbc.Config, tracer trace.Tracer) (_interface.Cache, error) {
	redisDB, err := dbc.OpenRedis(config)
	if err != nil {
		return nil, err
	}
	return &cache{
		redis:  redisDB,
		tracer: tracer,
	}, nil
}
