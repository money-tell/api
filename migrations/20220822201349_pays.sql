-- +goose Up
-- +goose StatementBegin
CREATE TYPE pays_repeat_type AS ENUM ('daily', 'weekly', 'monthly', 'yearly');
ALTER TYPE pays_repeat_type OWNER TO money;

CREATE TYPE pays_type AS ENUM ('accrual', 'redemption');
ALTER TYPE pays_type OWNER TO money;

CREATE TABLE pays
(
    id          BIGSERIAL    NOT NULL PRIMARY KEY,
    user_id     BIGINT       NOT NULL,
    type        pays_type    NOT NULL,
    title       VARCHAR(128) NOT NULL,
    amount      DECIMAL      NOT NULL,
    date        TIMESTAMP    NOT NULL,
    repeat_type pays_repeat_type DEFAULT NULL,
    created_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP,
    updated_at  TIMESTAMP        DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE pays;
DROP TYPE pays_repeat_type;
DROP TYPE pays_type;
-- +goose StatementEnd
