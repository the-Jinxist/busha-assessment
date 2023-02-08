-- name: CreateComment :one
INSERT INTO comments (
  movie_id, comment, comment_ip_address
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: ListComments :many
SELECT * FROM comments
WHERE movie_id = $1
ORDER BY date(created_at) DESC
LIMIT $2
OFFSET $3;