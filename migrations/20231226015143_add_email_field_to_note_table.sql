-- +goose Up
ALTER TABLE note ADD COLUMN email TEXT DEFAULT '';

-- +goose Down
ALTER TABLE note DROP COLUMN email;
