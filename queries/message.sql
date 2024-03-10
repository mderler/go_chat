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