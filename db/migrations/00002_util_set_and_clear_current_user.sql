-- +goose Up
-- +goose StatementBegin

-- 1. Create a schema for your helper functions:
CREATE SCHEMA IF NOT EXISTS util;

-- 2. Function to set the “current user” in the session:
CREATE OR REPLACE FUNCTION util.set_current_user_id(user_id INTEGER) RETURNS VOID AS $$
BEGIN
  PERFORM set_config('app.current_user_id', user_id::TEXT, FALSE);
END;
$$ LANGUAGE plpgsql;

-- 3. Function to clear it (e.g. on connection release):
CREATE OR REPLACE FUNCTION util.clear_user_id() RETURNS VOID AS $$
BEGIN
  PERFORM set_config('app.current_user_id', '-1', FALSE);
END;
$$ LANGUAGE plpgsql;

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- SELECT 'down SQL query';
-- +goose StatementEnd
