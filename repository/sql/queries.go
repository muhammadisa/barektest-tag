package sql

const (
	queryReadNewses = `SELECT id, topic_id, title, content, status, created_at, updated_at FROM news ORDER BY created_at DESC`
	queryWriteNews = `INSERT INTO news(id, topic_id, title, content, status, created_at, updated_at) VALUES (?,?,?,?,?,?,?)`
	queryUpdateNews = `UPDATE news SET topic_id = ?, title = ?, content = ?, status = ?, created_at = ?, updated_at = ? WHERE id = ?`
	queryRemoveNews = `DELETE FROM news WHERE id = ?`
	queryLookupCreateAtNews = `SELECT id, created_at FROM news WHERE id = ?`

	queryReadTags = `SELECT id, tag, created_at, updated_at FROM tags ORDER BY created_at DESC`
	queryWriteTag = `INSERT INTO tags(id, tag, created_at, updated_at) VALUES (?,?,?,?)`
	queryUpdateTag = `UPDATE tags SET tag = ?, created_at = ?, updated_at = ? WHERE id = ?`
	queryRemoveTag = `DELETE FROM tags WHERE id = ?`
	queryLookupCreateAtTag = `SELECT id, created_at FROM tags WHERE id = ?`

	queryReadTopics = `SELECT id, title, headline, created_at, updated_at FROM topics ORDER BY created_at DESC`
	queryWriteTopic = `INSERT INTO topics(id, title, headline, created_at, updated_at) VALUES (?,?,?,?,?)`
	queryUpdateTopic = `UPDATE topics SET title = ?, headline = ?, created_at = ?, updated_at = ? WHERE id = ?`
	queryRemoveTopic = `DELETE FROM topics WHERE id = ?`
	queryLookupCreateAtTopic = `SELECT id, created_at FROM topics WHERE id = ?`

)
