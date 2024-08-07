-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
    ALTER TABLE feed_follows ADD COLUMN tel_id INT NOT NULL;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
    ALTER TABLE feed_follows DROP COLUMN tel_id;