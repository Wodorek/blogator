-- name: CreateFeedFollow :one
WITH inserted_feed_follow AS (
INSERT INTO feed_follows
  (id, created_at, updated_at, user_id, feed_id)
VALUES
  ($1, $2, $3, $4, $5)
RETURNING *
)
SELECT
  inserted_feed_follow.*,
  feeds.name AS feed_name,
  users.name AS username
FROM inserted_feed_follow
  INNER JOIN feeds ON inserted_feed_follow.feed_id = feeds.id
  INNER JOIN users ON inserted_feed_follow.user_id = users.id;
--

-- name: GetFeedFollowsForUser :many
SELECT feed_follows.*, feeds.name AS feed_name, users.name AS username
FROM feed_follows
  INNER JOIN feeds ON feed_follows.feed_id = feeds.id
  INNER JOIN users ON feeds.user_id = users.id
WHERE feed_follows.user_id = $1;
--

-- name: DeleteFeedFollow :exec

WITH
  selected_feed
  AS
  (
    SELECT feeds.id AS selected_id
    FROM feeds
    WHERE feeds.URL = $1
  )
DELETE FROM feed_follows USING selected_feed WHERE feed_follows.user_id = $2 AND feed_follows.feed_id = selected_feed.selected_id;