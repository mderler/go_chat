-- name: CreateMessage :one
INSERT INTO message (
  sender_id, content
) VALUES (
  ?, ?
)
RETURNING id;

-- name: CreateDirectMessage :exec
INSERT INTO direct_message (
  receiver_id, message_id
) VALUES (
  ?, ?
);

-- name: GetDirectMessages :many
SELECT message.sender_id, message.content, user.full_name, user.color FROM message
JOIN user ON message.sender_id = user.id
JOIN direct_message ON message.id = direct_message.message_id
WHERE
  (direct_message.receiver_id = @user_id AND message.sender_id = @contact_id) OR 
  (direct_message.receiver_id = @contact_id AND message.sender_id = @user_id)
ORDER BY message.created_at ASC
LIMIT 10;