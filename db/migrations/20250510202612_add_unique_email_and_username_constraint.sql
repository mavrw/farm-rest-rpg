-- +goose Up
-- +goose StatementBegin
ALTER TABLE "user"
ADD CONSTRAINT unique_email UNIQUE (email);

ALTER TABLE "user"
ADD CONSTRAINT unique_username UNIQUE (username);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
