-- name: GetPaysByUserID :many
SELECT *
FROM pays
WHERE user_id = $1
  AND repeat_type is null;

-- name: GetRepeatedPaysByUserID :many
SELECT *
FROM pays
WHERE user_id = @user_id
  AND repeat_type is not null
  AND repeat_type is not null
  AND date between @from_date::timestamp and @to_date::timestamp
  AND (
            repeat_type = 'daily' OR
            (repeat_type = 'weekly' AND date_part('dow', date) in (@days_of_week::int[])) OR
            (repeat_type = 'monthly' AND date_part('day', date) between @monthly_from_day::int and @monthly_to_day::int) OR
            (repeat_type = 'yearly' AND date_part('doy', date) between @yearly_from_day::int and @yearly_to_day::int)
    );


-- name: PayInsert :one
INSERT INTO pays (user_id, type, title, amount, date, repeat_type)
VALUES (@user_id, @type, @title, @amount, @date, @repeat_type)
RETURNING id, user_id, type, title, amount, date, repeat_type, created_at, updated_at;