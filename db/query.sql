

-- name: CreateAuthor :one
INSERT INTO author (email, created_at, updated_at) VALUES (?, datetime('now'), datetime('now'))
RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM author WHERE id = ?;

-- name: ChangeAuthorEmail :exec
UPDATE author SET email = ?, updated_at = datetime('now') WHERE id = ?;


-- name: CreateEntry :one
INSERT INTO entry (title, body, author_id, created_at, updated_at) VALUES (?, ?, ?, datetime('now'), datetime('now'))
RETURNING *;

-- name: DeleteEntry :exec
DELETE FROM entry WHERE id = ?;

-- name: ChangeEntryTitle :exec
UPDATE entry SET title = ?, updated_at = datetime('now') WHERE id = ?;

-- name: ChangeEntryBody :exec
UPDATE entry SET body = ?, updated_at = datetime('now') WHERE id = ?;

-- name: GetEntryByAuthor :many
SELECT * FROM entry where author_id = ?;


-- name: CreateComment :one
INSERT INTO comment (body, author_id, entry_id, created_at, updated_at) VALUES (?, ? ,?, datetime('now'), datetime('now'))
RETURNING *;

-- name: DeleteComment :exec
DELETE FROM comment WHERE id = ?;

-- name: ChangeCommentBody :exec
UPDATE entry SET body = ?, updated_at = datetime('now') WHERE id = ?;

-- name: GetCommentsByEntry :many
SELECT * FROM comment where entry_id = ?;