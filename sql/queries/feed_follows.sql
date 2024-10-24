-- name: CreateFeedFollow :one

WITH inserted_feed_follow AS (
  INSERT INTO feed_follows (id, feed_id, user_id, created_at, updated_at)
  VALUES ($1, $2, $3, $4, $5)
  RETURNING *
)

SELECT ff.*, f.name AS feed_name, u.name AS user_name
FROM inserted_feed_follow AS ff
JOIN feeds AS f ON ff.feed_id = f.id
JOIN users AS u ON ff.user_id = u.id;



-- name: GetFeedFollowsForUser :many
SELECT ff.*, f.name AS feed_name, u.name AS user_name
FROM feed_follows AS ff
JOIN feeds AS f ON ff.feed_id = f.id
JOIN users AS u ON ff.user_id = u.idmake 
WHERE ff.user_id = $1;
