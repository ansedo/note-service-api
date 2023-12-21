-- +goose Up
CREATE TABLE note
(
    id         BIGSERIAL PRIMARY KEY,
    title      TEXT,
    text       TEXT,
    author     TEXT,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP
);

-- +goose Down
DROP TABLE note;
