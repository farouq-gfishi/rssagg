-- +goose Up
ALTER TABLE feeds ADD COLUMN fetched BOOLEAN DEFAULT FALSE;

-- +goose Down
ALTER TABLE feeds DROP COLUMN fetched;