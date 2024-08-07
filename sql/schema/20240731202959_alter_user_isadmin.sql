-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
    ALTER TABLE users ADD COLUMN is_admin VARCHAR(2) NOT NULL;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
    ALTER TABLE users DROP COLUMN is_admin;
