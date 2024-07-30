-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
    CREATE TABLE feeds (
        id UUID PRIMARY KEY NOT NULL,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP NOT NULL,
        name TEXT NOT NULL,
        slug TEXT UNIQUE NOT NULL,
        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE
    );

-- +goose Down
    DROP TABLE feeds;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
