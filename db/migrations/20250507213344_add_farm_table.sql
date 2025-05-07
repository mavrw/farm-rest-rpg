-- +goose Up
-- +goose StatementBegin
CREATE TABLE "farm" (
    "id"            SERIAL      PRIMARY KEY,
    "user_id"       INTEGER     REFERENCES "user"(id) NOT NULL,
    "name"          TEXT        NOT NULL,
    "created_at"    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at"    TIMESTAMP   NOT NULL DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "farm";
-- +goose StatementEnd
