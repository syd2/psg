// Code generated by sqlc. DO NOT EDIT.
// source: password.sql

package db

import (
	"context"
)

const createPassword = `-- name: CreatePassword :one

INSERT INTO passwords (app_name, app_password, user_id) 
VALUES ($1, $2, $3) RETURNING id, app_name, app_password, user_id
`

type CreatePasswordParams struct {
	AppName     string `json:"app_name"`
	AppPassword string `json:"app_password"`
	UserID      int64  `json:"user_id"`
}

func (q *Queries) CreatePassword(ctx context.Context, arg CreatePasswordParams) (Password, error) {
	row := q.db.QueryRowContext(ctx, createPassword, arg.AppName, arg.AppPassword, arg.UserID)
	var i Password
	err := row.Scan(
		&i.ID,
		&i.AppName,
		&i.AppPassword,
		&i.UserID,
	)
	return i, err
}

const deletePassword = `-- name: DeletePassword :exec

DELETE 
FROM passwords 
WHERE id = $1
`

func (q *Queries) DeletePassword(ctx context.Context, id int64) error {
	_, err := q.db.ExecContext(ctx, deletePassword, id)
	return err
}

const getPassword = `-- name: GetPassword :one

SELECT id, app_name, app_password, user_id 
FROM passwords
WHERE user_id = $1
AND app_name = $2
LIMIT 1
`

type GetPasswordParams struct {
	UserID  int64  `json:"user_id"`
	AppName string `json:"app_name"`
}

func (q *Queries) GetPassword(ctx context.Context, arg GetPasswordParams) (Password, error) {
	row := q.db.QueryRowContext(ctx, getPassword, arg.UserID, arg.AppName)
	var i Password
	err := row.Scan(
		&i.ID,
		&i.AppName,
		&i.AppPassword,
		&i.UserID,
	)
	return i, err
}

const updatePassword = `-- name: UpdatePassword :exec

UPDATE passwords
SET app_password = $1
WHERE user_id = $2
AND app_name = $3
`

type UpdatePasswordParams struct {
	AppPassword string `json:"app_password"`
	UserID      int64  `json:"user_id"`
	AppName     string `json:"app_name"`
}

func (q *Queries) UpdatePassword(ctx context.Context, arg UpdatePasswordParams) error {
	_, err := q.db.ExecContext(ctx, updatePassword, arg.AppPassword, arg.UserID, arg.AppName)
	return err
}