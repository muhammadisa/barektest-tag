package sql

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/muhammadisa/barektest-tag/model"
	pb "github.com/muhammadisa/barektest-tag/protoc/api/v1"
	_interface "github.com/muhammadisa/barektest-tag/repository/interface"
	"github.com/muhammadisa/barektest-util/dbc"
	uuid "github.com/satori/go.uuid"
	"go.opencensus.io/trace"
	"sync"
)

var mutex = &sync.RWMutex{}

type readWrite struct {
	tracer trace.Tracer
	db     *sql.DB
}

func (r *readWrite) WriteTag(ctx context.Context, req *pb.Tag) (err error) {
	const query = `
	INSERT INTO tags(id, tag) VALUES (?,?)
	`
	req.Id = uuid.NewV4().String()
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(
		ctx,
		req.Id,  // id
		req.Tag, // tag
	)
	if err != nil {
		return err
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return fmt.Errorf("failed to insert reason : %+v", err)
	}
	return nil
}

func (r *readWrite) ModifyTag(ctx context.Context, req *pb.Tag) (err error) {
	const query = `
	UPDATE tags SET tag = ? WHERE id = ?
	`
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmt.ExecContext(
		ctx,
		req.Tag, // tag
		req.Id,  // id
	)
	if err != nil {
		return err
	}
	if affected, err := result.RowsAffected(); affected == 0 || err != nil {
		return fmt.Errorf("failed to insert reason : %+v", err)
	}
	return nil
}

func (r *readWrite) RemoveTag(ctx context.Context, req *pb.SelectTag) (err error) {
	const query = `
	DELETE FROM tags WHERE id = ?
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

func (r *readWrite) ReadTags(ctx context.Context) (res *pb.Tags, err error) {
	var tags pb.Tags
	var tag model.Tag
	const query = `
	SELECT id, tag, created_at, updated_at FROM tags ORDER BY created_at DESC 
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
			&tag.ID,      // id
			&tag.Tag,     // tag
			&tag.Created, // created_at
			&tag.Updated, // updated_at
		)
		if err != nil {
			return res, err
		}
		tag.UseUnixTimeStamp()
		tags.Tags = append(tags.Tags, &pb.Tag{
			Id:        tag.ID,
			Tag:       tag.Tag,
			CreatedAt: tag.CreatedAt,
			UpdatedAt: tag.UpdatedAt,
		})
	}
	return &tags, nil
}

func NewSQL(config dbc.Config, tracer trace.Tracer) (_interface.ReadWrite, error) {
	sqlDB, err := dbc.OpenDB(config)
	if err != nil {
		return nil, err
	}
	return &readWrite{
		db:     sqlDB,
		tracer: tracer,
	}, nil
}
