-- +goose Up
-- +goose StatementBegin
ALTER TABLE user ADD COLUMN color VARCHAR(7) NOT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE user DROP COLUMN color;
-- +goose StatementEnd
