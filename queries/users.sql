-- name: FindUser :one
SELECT *
FROM users
WHERE email = $1
  AND password = $2
  AND is_active = true;