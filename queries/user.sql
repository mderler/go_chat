-- name: CreateUser :execrows
INSERT INTO user (
  username, password
) VALUES (
  ?, ?
);

-- name: ListUser :many
SELECT id, username FROM user
ORDER BY name;

-- name: GetIdAndPasswordByUsername :one
SELECT id, password FROM user
WHERE username = ?;