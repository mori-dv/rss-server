-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
    ALTER TABLE users ADD COLUMN tel_id INT;
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
    ALTER TABLE users DROP COLUMN tel_id;
