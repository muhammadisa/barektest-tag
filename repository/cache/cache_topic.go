package cache

import (
	"context"
	"encoding/json"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
)

func (c *cache) ReloadTopics(ctx context.Context, topics *pb.Topics) (err error) {
	data := make(map[string]interface{})
	for _, topic := range topics.Topics {
		tagByte, err := json.Marshal(topic)
		if err != nil {
			return err
		}
		data[topic.Id] = string(tagByte)
	}
	return c.redis.HMSet(ctx, "topics", data).Err()
}

func (c *cache) GetTopics(ctx context.Context) (res *pb.Topics, err error) {
	var topics pb.Topics
	tagsMap := c.redis.HGetAll(ctx, "topics").Val()
	for _, v := range tagsMap {
		var topic pb.Topic
		err = json.Unmarshal([]byte(v), &topic)
		if err != nil {
			return res, err
		}
		topics.Topics = append(topics.Topics, &topic)
	}
	return &topics, nil
}

func (c *cache) UnsetTopic(ctx context.Context, id string) (err error) {
	return c.redis.HDel(ctx, "topics", id).Err()
}

func (c *cache) SetTopic(ctx context.Context, topic *pb.Topic) (err error) {
	tagByte, err := json.Marshal(topic)
	if err != nil {
		return err
	}
	return c.redis.HSetNX(ctx, "topics", topic.Id, string(tagByte)).Err()
}
