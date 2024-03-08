-- +goose Up
-- +goose StatementBegin
CREATE TABLE message (
    id SERIAL PRIMARY KEY,
    sender_id INTEGER NOT NULL,
    message VARCHAR(1024) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (sender_id) REFERENCES users (id)
);

CREATE TABLE direct_message (
    receiver_id INTEGER NOT NULL,
    message_id INTEGER NOT NULL UNIQUE,
    FOREIGN KEY (receiver_id) REFERENCES users (id),
    FOREIGN KEY (message_id) REFERENCES message (id),
    PRIMARY KEY (receiver_id, message_id)
);

CREATE TABLE "group" (
    id SERIAL PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE user_group (
    user_id INTEGER NOT NULL,
    group_id INTEGER NOT NULL,
    FOREIGN KEY (user_id) REFERENCES users (id),
    FOREIGN KEY (group_id) REFERENCES "group" (id),
    PRIMARY KEY (user_id, group_id)
);

CREATE TABLE group_message (
    group_id INTEGER NOT NULL,
    message_id INTEGER NOT NULL UNIQUE,
    FOREIGN KEY (group_id) REFERENCES "group" (id),
    FOREIGN KEY (message_id) REFERENCES message (id),
    PRIMARY KEY (group_id, message_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE message;
DROP TABLE direct_message;
DROP TABLE group;
DROP TABLE user_group;
DROP TABLE group_message;
-- +goose StatementEnd
