-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name, api_key)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex') )
RETURNING *;

-- name: GetUser :one
SELECT * From users where api_key = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: DelUser :execresult
DELETE FROM users WHERE id = $1;
