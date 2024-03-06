-- name: CreateUser :one
INSERT INTO user (
  username, full_name, password
) VALUES (
  ?, ?, ?
) RETURNING id;

-- name: ListUser :many
SELECT id, username, full_name FROM user
ORDER BY name;

-- name: GetUserForLogin :one
SELECT id, password FROM user
WHERE username = ?;

-- name: GetFullNameById :one
SELECT full_name FROM user
WHERE id = ?;