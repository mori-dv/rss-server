-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
    CREATE TABLE feed_follows (
        id UUID PRIMARY KEY,
        created_at TIMESTAMP NOT NULL,
        updated_at TIMESTAMP Not NULL,
        user_id UUID NOT NULL REFERENCES users(id) ON DELETE CASCADE,
        feed_id UUID NOT NULL REFERENCES feeds(id) ON DELETE CASCADE,
        UNIQUE(user_id, feed_id)
    );
-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
DROP TABLE feed_follows;
