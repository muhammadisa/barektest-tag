package model

import "time"

type News struct {
	ID               string
	TopicID          string
	Title            string
	Content          string
	Status           int32
	CreatedAt        int64
	UpdatedAt        int64
	Created, Updated time.Time
}

func (news *News) UseUnixTimeStamp() {
	news.CreatedAt = news.Created.Unix()
	news.UpdatedAt = news.Updated.Unix()
}
