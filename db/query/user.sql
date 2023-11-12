-- name: CreateUser :one
INSERT INTO users (
    username,
    hashed_password,
    full_name,
    email
) VALUES
    ($1, $2, $3, $4)
RETURNING *;

-- name: GetUser :one
SELECT * FROM users WHERE username = $1 LIMIT 1;

-- name: ChangePassword :one
UPDATE users SET hashed_password = $2 password_updated_at = NOW() WHERE username = $1 RETURNING *;

-- name: UpdateUser :one
UPDATE users SET
    full_name = $2,
    email = $3
WHERE username = $1