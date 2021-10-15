package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/muhammadisa/barektest-tag/model"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
	uuid "github.com/satori/go.uuid"
	"time"
)

func (r *readWrite) WriteTopic(ctx context.Context, req *pb.Topic) (res *pb.Topic, err error) {
	const query = `
	INSERT INTO topics(id, title, headline, created_at, updated_at) VALUES (?,?,?,?,?)
	`
	currentTime := time.Now()
	req.Id = uuid.NewV4().String()
	req.CreatedAt = currentTime.Unix()
	req.UpdatedAt = currentTime.Unix()
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return res, err
	}
	result, err := stmt.ExecContext(
		ctx,
		req.Id,       // id
		req.Title,    // title
		req.Headline, // headline
		currentTime,  // created_at
		currentTime,  // updated_at
	)
	if err != nil {
		return res, err
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return res, fmt.Errorf("failed to insert reason : %+v", err)
	}
	return req, nil
}

func (r *readWrite) ModifyTopic(ctx context.Context, req *pb.Topic) (res *pb.Topic, err error) {
	var oldTopic model.Topic
	const queryLookupExisting = `
	SELECT id, title, headline, created_at, updated_at FROM topics WHERE id = ? 
	`
	stmt, err := r.db.Prepare(queryLookupExisting)
	if err != nil {
		return res, err
	}
	row := stmt.QueryRow(req.Id)
	err = row.Scan(
		&oldTopic.ID,       // id
		&oldTopic.Title,    // title
		&oldTopic.Headline, // headline
		&oldTopic.Created,  // created_at
		&oldTopic.Updated,  // updated_at
	)
	if err != nil || err == sql.ErrNoRows {
		return res, err
	}
	oldTopic.UseUnixTimeStamp()

	const query = `
	UPDATE topics SET title = ?, headline = ?, created_at = ?, updated_at = ? WHERE id = ?
	`
	currentTime := time.Now()
	req.CreatedAt = oldTopic.CreatedAt
	req.UpdatedAt = currentTime.Unix()
	stmt, err = r.db.Prepare(query)
	if err != nil {
		return res, err
	}
	result, err := stmt.ExecContext(
		ctx,
		req.Title,        // title
		req.Headline,     // headline
		oldTopic.Created, // created_at
		currentTime,      // updated_at
		req.Id,           // id
	)
	if err != nil {
		return res, err
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return res, fmt.Errorf("failed to insert reason : %+v", err)
	}
	return req, nil
}

func (r *readWrite) RemoveTopic(ctx context.Context, req *pb.Select) error {
	const query = `
	DELETE FROM topics WHERE id = ?
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(
		ctx,
		req.Id, // id
	)
	if err != nil {
		return err
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return fmt.Errorf("failed to insert reason : %+v", err)
	}
	return nil
}

func (r *readWrite) ReadTopics(ctx context.Context) (res *pb.Topics, err error) {
	var topics pb.Topics
	var topic model.Topic
	const query = `
	SELECT id, title, headline, created_at, updated_at FROM topics ORDER BY created_at DESC 
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return res, err
	}
	mutex.Lock()
	row, err := stmt.QueryContext(ctx)
	if err != nil {
		return res, err
	}
	mutex.Unlock()
	for row.Next() {
		err = row.Scan(
			&topic.ID,       // id
			&topic.Title,    // title
			&topic.Headline, // headline
			&topic.Created,  // created_at
			&topic.Updated,  // updated_at
		)
		if err != nil {
			return res, err
		}
		topic.UseUnixTimeStamp()
		topics.Topics = append(topics.Topics, &pb.Topic{
			Id:        topic.ID,
			Title:     topic.Title,
			Headline:  topic.Headline,
			CreatedAt: topic.CreatedAt,
			UpdatedAt: topic.UpdatedAt,
		})
	}
	return &topics, nil
}
