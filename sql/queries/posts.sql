-- name: CreatePost :exec
INSERT INTO posts
  (id,created_at,updated_at,title,url,description,published_at,feed_id)
VALUES
  ($1, $2, $3, $4, $5, $6, $7, $8);

-- name: GetPostsForUser :many
SELECT *
FROM posts
  INNER JOIN feed_follows ON posts.feed_id = feed_follows.feed_id
  INNER JOIN users ON users.id = feed_follows.user_id
WHERE users.id = $1
ORDER BY posts.published_at DESC
LIMIT $2;
