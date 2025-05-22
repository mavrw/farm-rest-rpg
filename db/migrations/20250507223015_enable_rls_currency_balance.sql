-- +goose Up
-- +goose StatementBegin

ALTER TABLE currency_balance ENABLE ROW LEVEL SECURITY;

CREATE POLICY currency_balance_isolation_policy
  ON currency_balance
  USING (
    current_setting('app.current_user_id')::INTEGER = user_id
  );

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
