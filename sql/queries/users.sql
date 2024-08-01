-- name: CreateUser :one
INSERT INTO users(id, created_at, updated_at, name, api_key, is_admin)
VALUES ($1, $2, $3, $4, encode(sha256(random()::text::bytea), 'hex'), $5)
RETURNING *;

-- name: GetUser :one
SELECT * From users where api_key = $1;

-- name: GetAllUsers :many
SELECT * FROM users;

-- name: GetAdminUsers :many
SELECT * FROM users WHERE is_admin=1;

-- name: CheckAdmin :one
SELECT name, api_key, is_admin FROM users WHERE id=$1;

-- name: UpdateUser: