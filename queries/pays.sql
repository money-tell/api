-- name: GetPaysByUserID :many
SELECT id,
       user_id,
       type,
       title,
       amount,
       date,
       repeat_type,
       created_at,
       updated_at
FROM pays
WHERE user_id = @user_id::uuid
  AND date between @date_from::timestamp and @date_to::timestamp
  AND repeat_type = 'none';

-- name: GetRepeatedPaysByUserID :many
SELECT id,
       user_id,
       type,
       title,
       amount,
       date,
       repeat_type,
       created_at,
       updated_at
FROM pays
WHERE user_id = @user_id
  AND repeat_type != 'none'
  AND (
            repeat_type = 'daily' OR
            (repeat_type = 'weekly' AND date_part('dow', date) = ANY(@days_of_week::int[])) OR
            (repeat_type = 'monthly' AND
             date_part('day', date) between @monthly_day_from::int and @monthly_day_to::int) OR
            (repeat_type = 'yearly' AND date_part('doy', date) between @yearly_day_from::int and @yearly_day_to::int)
    );


-- name: PayInsert :one
INSERT INTO pays (user_id, type, title, amount, date, repeat_type)
VALUES (@user_id, @type, @title, @amount, @date, @repeat_type)
RETURNING id, user_id, type, title, amount, date, repeat_type, created_at, updated_at;