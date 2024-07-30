-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

    CREATE TABLE users (
        id UUID PRIMARY KEY NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        name TEXT NOT NULL
    );
-- +goose Down
    DROP TABLE users;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
