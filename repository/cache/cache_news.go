package cache

import (
	"context"
	"encoding/json"
	"github.com/muhammadisa/barektest-tag/constant"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
)

func (c *cache) GetNewses(ctx context.Context) (res *pb.Newses, err error) {
	var newses pb.Newses
	tagsMap := c.redis.HGetAll(ctx, constant.News).Val()
	for _, v := range tagsMap {
		var news pb.News
		err = json.Unmarshal([]byte(v), &news)
		if err != nil {
			return res, err
		}
		newses.Newses = append(newses.Newses, &news)
	}
	c.redis.Set(ctx, constant.ReloadNewses, false, 0)
	return &newses, nil
}

func (c *cache) UnsetNews(ctx context.Context, id string) error {
	return c.redis.HDel(ctx, constant.News, id).Err()
}

func (c *cache) InvalidateNewses(ctx context.Context) error {
	return c.redis.Set(ctx, constant.ReloadNewses, true, 0).Err()
}

func (c *cache) ReloadRequired(ctx context.Context) bool {
	state := c.redis.Get(ctx, constant.ReloadNewses).Val()
	if state != "0" {
		return true
	} else {
		return false
	}
}

func (c *cache) ReloadNewses(ctx context.Context, newses *pb.Newses) (err error) {
	data := make(map[string]interface{})
	for _, news := range newses.Newses {
		newsByte, err := json.Marshal(news)
		if err != nil {
			return err
		}
		data[news.Id] = string(newsByte)
	}
	c.redis.Set(ctx, constant.ReloadNewses, false, 0)
	return c.redis.HMSet(ctx, constant.News, data).Err()
}
