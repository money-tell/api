-- +goose Up
-- +goose StatementBegin
CREATE
EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE transactions_repeat_type AS ENUM ('daily', 'weekly', 'monthly', 'yearly', 'none');
ALTER TYPE transactions_repeat_type OWNER TO money;

CREATE TYPE transactions_type AS ENUM ('accrual', 'redemption');
ALTER TYPE transactions_type OWNER TO money;

-- transactions.sql
CREATE TABLE transactions
(
    id          UUID                              DEFAULT uuid_generate_v4() NOT NULL PRIMARY KEY,
    user_id     UUID                     NOT NULL,
    type        transactions_type        NOT NULL,
    title       VARCHAR(128)             NOT NULL,
    amount      DECIMAL                  NOT NULL,
    date        TIMESTAMP                NOT NULL,
    repeat_type transactions_repeat_type NOT NULL DEFAULT 'none',
    created_at  TIMESTAMP                         DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP                         DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE transactions;
DROP TYPE transactions_repeat_type;
DROP TYPE transactions_type;
-- +goose StatementEnd
