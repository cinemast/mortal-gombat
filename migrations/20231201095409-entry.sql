
-- +migrate Up
CREATE TABLE entry(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    title VARCHAR(180) NOT NULL,
    body TEXT,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    author_id INTEGER NOT NULL,
    FOREIGN KEY (author_id) REFERENCES author(id)
);
-- +migrate Down
DROP TABLE entry;