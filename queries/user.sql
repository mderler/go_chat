-- name: CreateUser :one
INSERT INTO user (
  username, full_name, password, color
) VALUES (
  ?, ?, ?, ?
) RETURNING id;

-- name: ListUser :many
SELECT id, username, full_name FROM user
ORDER BY name;

-- name: GetUserForLogin :one
SELECT id, password FROM user
WHERE username = ?;

-- name: GetUserForChatById :one
SELECT full_name, color FROM user
WHERE id = ?;

-- name: GetUsersByQuery :many
SELECT id, username, full_name FROM user
WHERE (username LIKE @name OR full_name LIKE @name) AND id != @id;

-- name: GetContactedUsers :many
SELECT user.id, user.full_name, user.color FROM user
JOIN direct_message
JOIN message ON message.id = direct_message.message_id
WHERE
  (direct_message.receiver_id = @user_id AND message.sender_id = user.id) OR 
  (direct_message.receiver_id = user.id AND message.sender_id = @user_id) AND
  user.id != @user_id
GROUP BY user.id
ORDER BY message.created_at DESC;

-- name: GetLastContactedUser :one
SELECT user.id, user.full_name, user.color FROM user
JOIN direct_message
JOIN message ON message.id = direct_message.message_id
WHERE
  (direct_message.receiver_id = @user_id AND message.sender_id = user.id) OR 
  (direct_message.receiver_id = user.id AND message.sender_id = @user_id) AND
  user.id != @user_id
ORDER BY message.created_at DESC
LIMIT 1;