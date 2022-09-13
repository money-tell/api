-- name: GetPaysByUserID :many
SELECT *
FROM pays
WHERE user_id = $1;

-- name: PayInsert :one
INSERT INTO  pays (user_id, type, title, amount, date, repeat_type)
VALUES (@user_id, @type, @title, @amount, @date, @repeat_type)
RETURNING id, user_id, type, title, amount, date, repeat_type, created_at, updated_at;