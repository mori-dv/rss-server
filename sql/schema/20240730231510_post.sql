-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    title VARCHAR(250) NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    published_at TIMESTAMP NOT NULL,
    url VARCHAR(250) UNIQUE NOT NULL,
    description TEXT,
    feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE
);
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE posts;