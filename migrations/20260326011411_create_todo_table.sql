-- +goose Up
CREATE TABLE IF NOT EXISTS todo (
    id         TEXT    PRIMARY KEY,
    title      TEXT    NOT NULL,
    done       INTEGER NOT NULL DEFAULT 0,
    created_at TEXT    NOT NULL DEFAULT (datetime('now'))
);

-- +goose Down
DROP TABLE IF EXISTS todo;
