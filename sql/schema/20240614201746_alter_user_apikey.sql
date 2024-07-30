-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

    Alter TABLE users Add COLUMN api_key VARCHAR(64) UNIQUE NOT NULL DEFAULT (
        encode(sha256(random()::text::bytea), 'hex')
    );

-- +goose Down
    ALTER TABLE users DROP COLUMN api_key;
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd
