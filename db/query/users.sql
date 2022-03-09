
-- name: GetUser :one
SELECT * FROM users
WHERE username = $1 LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1;

-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES ($1, $2) RETURNING *;

-- name: DeleteUser :exec
DELETE FROM users WHERE id = $1;