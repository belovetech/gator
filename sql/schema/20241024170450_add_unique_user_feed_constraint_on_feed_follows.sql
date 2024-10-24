-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
ALTER TABLE feed_follows ADD CONSTRAINT unique_feed_user UNIQUE (feed_id, user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
ALTER TABLE feed_follows DROP CONSTRAINT unique_feed_user;
-- +goose StatementEnd

