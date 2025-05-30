package models

import "database/sql"

// Issue represents an issue.
type Issue struct {
	ID        int64         `db:"id"`
	RepoID    int64         `db:"repo_id"`
	Assignee  sql.NullInt64 `db:"assignee_id"`
	Title     string        `db:"title"`
	Body      string        `db:"body"`
	CreatedAt string        `db:"created_at"`
	UpdatedAt string        `db:"updated_at"`
}
