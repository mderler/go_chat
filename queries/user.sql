-- name: CreateUser :execrows
INSERT INTO user (
  username, password
) VALUES (
  ?, ?
);

-- name: ListUser :many
SELECT id, username FROM user
ORDER BY name;

-- name: GetPasswordByUsername :one
SELECT password FROM user
WHERE username = ?;