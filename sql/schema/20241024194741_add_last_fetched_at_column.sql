-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE feeds ADD COLUMN last_fetched_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE feeds DROP COLUMN last_fetched_at;
-- +goose StatementEnd
