

-- name: CreateAuthor :one
INSERT INTO author (email, created_at, updated_at) VALUES (?, datetime('now'), datetime('now'))
RETURNING *;

-- name: CreateEntry :one
INSERT INTO entry (title, body, author_id, created_at, updated_at) VALUES (?, ?, ?, datetime('now'), datetime('now'))
RETURNING *;

-- name: CreateComment :one
INSERT INTO comment (body, author_id, entry_id, created_at, updated_at) VALUES (?, ? ,?, datetime('now'), datetime('now'))
RETURNING *;