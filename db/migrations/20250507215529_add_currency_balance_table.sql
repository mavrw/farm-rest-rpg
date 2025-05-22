-- +goose Up
-- +goose StatementBegin
CREATE TABLE "currency_balance" (
    "id"                    SERIAL                                      PRIMARY KEY,
    "user_id"               INTEGER REFERENCES "users"(id)              NOT NULL,
    "currency_type_id"      INTEGER REFERENCES "currency_type"(id)      NOT NULL,
    "balance"               INTEGER                                     NOT NULL DEFAULT 0 CHECK (balance >= 0),

    -- prevent duplicate currency rows per user
    UNIQUE ("user_id", "currency_type_id")
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE "currency_balance";
-- +goose StatementEnd
