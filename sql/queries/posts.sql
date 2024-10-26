-- name: CreatePost :one
INSERT INTO posts (id, feed_id, title, url, description, created_at, updated_at, published_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING *;


-- name: GetPostByUser :many
SELECT p.*, f.*, u.*
FROM posts AS p
JOIN feeds as f ON p.feed_id = f.id
JOIN users as u ON f.user_id = u.id
WHERE u.name = ($1)
ORDER BY p.published_at DESC
LIMIT $2;
