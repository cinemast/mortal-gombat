
-- +migrate Up
CREATE TABLE comment(
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    body TEXT NOT NULL,
    entry_id INTEGER NOT NULL,
    author_id INTEGER NOT NULL,
    created_at TEXT NOT NULL,
    updated_at TEXT NOT NULL,
    FOREIGN KEY (entry_id) REFERENCES entry(id),
    FOREIGN KEY (author_id) REFERENCES author(id)
);
-- +migrate Down
DROP TABLE comment;