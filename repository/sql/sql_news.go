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

func (r *readWrite) WriteNews(ctx context.Context, req *pb.News) (res *pb.News, err error) {
	const query = `
	INSERT INTO news(id, topic_id, title, content, status, created_at, updated_at) VALUES (?,?,?,?,?,?,?)
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
		req.Id,      // id
		req.TopicId, // topic_id
		req.Title,   // title
		req.Content, // content
		req.Status,  // status
		currentTime, // created_at
		currentTime, // updated_at
	)
	if err != nil {
		return res, err
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return res, fmt.Errorf("failed to insert reason : %+v", err)
	}
	return req, nil
}

func (r *readWrite) ModifyNews(ctx context.Context, req *pb.News) (res *pb.News, err error) {
	var oldNews model.News
	const queryLookupExisting = `
	SELECT id, topic_id, title, content, status, created_at, updated_at FROM news WHERE id = ? 
	`
	stmt, err := r.db.Prepare(queryLookupExisting)
	if err != nil {
		return res, err
	}
	row := stmt.QueryRow(req.Id)
	err = row.Scan(
		&oldNews.ID,      // id
		&oldNews.TopicID, // topic_id
		&oldNews.Title,   // title
		&oldNews.Content, // content
		&oldNews.Status,  // status
		&oldNews.Created, // created_at
		&oldNews.Updated, // updated_at
	)
	if err != nil || err == sql.ErrNoRows {
		return res, err
	}
	oldNews.UseUnixTimeStamp()

	const query = `
	UPDATE news SET topic_id = ?, title = ?, content = ?, status = ?, created_at = ?, updated_at = ? WHERE id = ?
	`
	currentTime := time.Now()
	req.CreatedAt = oldNews.CreatedAt
	req.UpdatedAt = currentTime.Unix()
	stmt, err = r.db.Prepare(query)
	if err != nil {
		return res, err
	}
	result, err := stmt.ExecContext(
		ctx,
		req.TopicId,     // topic_id
		req.Title,       // title
		req.Content,     // content
		req.Status,      // status
		oldNews.Created, // created_at
		currentTime,     // updated_at
		req.Id,          // id
	)
	if err != nil {
		return res, err
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return res, fmt.Errorf("failed to insert reason : %+v", err)
	}
	return req, nil
}

func (r *readWrite) RemoveNews(ctx context.Context, req *pb.Select) error {
	const query = `
	DELETE FROM news WHERE id = ?
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

func (r *readWrite) ReadNewses(ctx context.Context) (res *pb.Newses, err error) {
	var newses pb.Newses
	var news model.News
	const query = `
	SELECT id, topic_id, title, content, status, created_at, updated_at FROM news ORDER BY created_at DESC 
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
			&news.ID,      // id
			&news.TopicID, // topic_id
			&news.Title,   // title
			&news.Content, // content
			&news.Status,  // status
			&news.Created, // created_at
			&news.Updated, // updated_at
		)
		if err != nil {
			return res, err
		}
		news.UseUnixTimeStamp()
		newses.Newses = append(newses.Newses, &pb.News{
			Id:        news.ID,
			TopicId:   news.TopicID,
			Title:     news.Title,
			Content:   news.Content,
			Status:    news.Status,
			CreatedAt: news.CreatedAt,
			UpdatedAt: news.UpdatedAt,
		})
	}
	return &newses, nil
}
