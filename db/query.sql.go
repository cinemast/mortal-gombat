// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0
// source: query.sql

package db

import (
	"context"
	"database/sql"
)

const createAuthor = `-- name: CreateAuthor :one
INSERT INTO author (email, created_at, updated_at) VALUES (?, datetime('now'), datetime('now'))
RETURNING id, email, created_at, updated_at
`

func (q *Queries) CreateAuthor(ctx context.Context, email string) (Author, error) {
	row := q.db.QueryRowContext(ctx, createAuthor, email)
	var i Author
	err := row.Scan(
		&i.ID,
		&i.Email,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createComment = `-- name: CreateComment :one
INSERT INTO comment (body, author_id, entry_id, created_at, updated_at) VALUES (?, ? ,?, datetime('now'), datetime('now'))
RETURNING id, body, entry_id, author_id, created_at, updated_at
`

type CreateCommentParams struct {
	Body     string
	AuthorID int64
	EntryID  int64
}

func (q *Queries) CreateComment(ctx context.Context, arg CreateCommentParams) (Comment, error) {
	row := q.db.QueryRowContext(ctx, createComment, arg.Body, arg.AuthorID, arg.EntryID)
	var i Comment
	err := row.Scan(
		&i.ID,
		&i.Body,
		&i.EntryID,
		&i.AuthorID,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const createEntry = `-- name: CreateEntry :one
INSERT INTO entry (title, body, author_id, created_at, updated_at) VALUES (?, ?, ?, datetime('now'), datetime('now'))
RETURNING id, title, body, created_at, updated_at, author_id
`

type CreateEntryParams struct {
	Title    string
	Body     sql.NullString
	AuthorID int64
}

func (q *Queries) CreateEntry(ctx context.Context, arg CreateEntryParams) (Entry, error) {
	row := q.db.QueryRowContext(ctx, createEntry, arg.Title, arg.Body, arg.AuthorID)
	var i Entry
	err := row.Scan(
		&i.ID,
		&i.Title,
		&i.Body,
		&i.CreatedAt,
		&i.UpdatedAt,
		&i.AuthorID,
	)
	return i, err
}