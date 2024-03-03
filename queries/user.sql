-- name: CreateUser :one
INSERT INTO user (
  username, password
) VALUES (
  ?, ?
)
RETURNING id, username;

-- name: ListUser :many
SELECT id, username FROM user
ORDER BY name;