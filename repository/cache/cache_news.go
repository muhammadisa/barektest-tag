package cache

import (
	"context"
	"encoding/json"
	"github.com/muhammadisa/barektest-tag/constant"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
)

func (c *cache) SetNews(ctx context.Context, news *pb.News) (err error) {
	newsByte, err := json.Marshal(news)
	if err != nil {
		return err
	}
	return c.redis.HSetNX(ctx, constant.News, news.Id, string(newsByte)).Err()
}

func (c *cache) UnsetNews(ctx context.Context, id string) (err error) {
	return c.redis.HDel(ctx, constant.News, id).Err()
}

func (c *cache) GetNewses(ctx context.Context) (res *pb.Newses, err error) {
	var newses pb.Newses
	newsesMap := c.redis.HGetAll(ctx, constant.News).Val()
	for _, v := range newsesMap {
		var news pb.News
		err = json.Unmarshal([]byte(v), &news)
		if err != nil {
			return res, err
		}
		newses.Newses = append(newses.Newses, &news)
	}
	return &newses, nil
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
	return c.redis.HMSet(ctx, constant.News, data).Err()
}

