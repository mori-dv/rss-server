-- name: CreateFeedFollows :one
INSERT INTO feed_follows (id, created_at, updated_at, user_id, feed_id, tel_id)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- name: GetFeedFollows :many
SELECT * FROM feed_follows WHERE user_id=$1;

-- name: DeleteFeedFollows :exec
DELETE FROM feed_follows WHERE feed_id=$1 AND user_id=$2;

-- name: GetFeedFollowsByTelId :many
SELECT * FROM feed_follows WHERE tel_id=$1;