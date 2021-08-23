
-- name: GetUser :one
SELECT * FROM users
WHERE id = $1 LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES ($1, $2) RETURNING *;

-- name: GetUserName :one
SELECT username FROM users
WHERE id = $1 LIMIT 1;

--name: DeleteUser :exec
DELETE FROM users WHERE id = $1;