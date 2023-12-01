
-- +migrate Up
CREATE TABLE author(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    email VARCHAR(100) NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL
);
-- +migrate Down
DROP TABLE author;