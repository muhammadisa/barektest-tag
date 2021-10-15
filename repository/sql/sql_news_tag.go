package sql

import (
	"context"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strings"
	"time"
)

func (r *readWrite) ReadNewsTagsTagIDAndTagByNewsID(ctx context.Context, newsID string, all bool) (res []string) {
	var tagID, tag string
	stmt, err := r.db.Prepare(queryReadNewsTags)
	if err != nil {
		return nil
	}
	mutex.Lock()
	row, err := stmt.QueryContext(ctx, newsID)
	if err != nil {
		return nil
	}
	mutex.Unlock()
	for row.Next() {
		err = row.Scan(
			&tagID,
			&tag,
		)
		if err != nil {
			return nil
		}
		if all {
			res = append(res, fmt.Sprintf("%s:%s", tagID, tag))
		}else{
			res = append(res, tagID)
		}
	}
	return res
}

func (r *readWrite) WriteNewsTags(ctx context.Context, newsID string, tagIDs []string) error {
	timeNow := time.Now()
	length := len(tagIDs)

	if length < 0 {
		return nil
	}

	var valueStrings []string
	var valueArgs []interface{}
	for _, tagID := range tagIDs {
		id := uuid.NewV4().String()
		valueStrings = append(valueStrings, "(?,?,?,?,?)")
		valueArgs = append(valueArgs, id)
		valueArgs = append(valueArgs, newsID)
		valueArgs = append(valueArgs, tagID)
		valueArgs = append(valueArgs, timeNow)
		valueArgs = append(valueArgs, timeNow)
	}

	query := fmt.Sprintf(queryWriteBulkNewsTags, strings.Join(valueStrings, ","))
	fmt.Println(query)
	result, err := r.db.Exec(query, valueArgs...)
	if err != nil {
		return err
	}
	if affected, err := result.RowsAffected(); affected != int64(length) || err != nil {
		return fmt.Errorf("failed to insert reason : %+v", err)
	}
	return nil
}
