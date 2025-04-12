-- name: CreateFeed :one
INSERT INTO feeds (id, name, url, user_id, created_at, updated_at)
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;

-- -- name: GetUserFeeds :many
-- SELECT * FROM feeds
-- WHERE user_id = $1

-- -- name: GetFeedByID :one
-- SELECT * FROM feeds
-- WHERE id = $1