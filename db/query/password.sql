-- name : CreatePassword :exec

INSERT INTO passwords (app_name, app_password, user_id) 
VALUES ($1, $2, $3) RETURNING *;

-- name: GetPassword :one

SELECT * 
FROM passwords
WHERE user_id = $1
AND app_name = $2
LIMIT 1;

-- name: DeletePassword :exec

DELETE 
FROM passwords 
WHERE id = $1;

-- name: UpdatePassword :exec

UPDATE passwords
SET app_password = $1
WHERE user_id = $2
AND app_name = $3;