-- name: CreateFeed :one

INSERT INTO feeds (id, name, url, user_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;


-- name: GetFeed :one
SELECT * FROM feeds WHERE url = ($1);


-- name: GetFeeds :many
SELECT f.name, f.url, u.name AS user_name FROM feeds AS f
JOIN users AS u ON f.user_id = u.id;

