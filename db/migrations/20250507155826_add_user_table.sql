-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user" (
    "id"            SERIAL      PRIMARY KEY,
    "username"      TEXT        NOT NULL,
    "email"         TEXT        DEFAULT NULL,
    "password_hash" TEXT        NOT NULL,
    "created_at"    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "user";
-- +goose StatementEnd
