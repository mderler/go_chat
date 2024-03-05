-- name: CreateUser :execrows
INSERT INTO user (
  username, full_name, password
) VALUES (
  ?, ?, ?
);

-- name: ListUser :many
SELECT id, username, full_name FROM user
ORDER BY name;

-- name: GetPasswordByUsername :one
SELECT password FROM user
WHERE username = ?;