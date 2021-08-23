// Code generated by sqlc. DO NOT EDIT.
// source: users.sql

package db

import (
	"context"
)

const createUser = `-- name: CreateUser :one
INSERT INTO users (username, password)
VALUES ($1, $2) RETURNING id, username, password
`

type CreateUserParams struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRowContext(ctx, createUser, arg.Username, arg.Password)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT id, username, password FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, id int64) (User, error) {
	row := q.db.QueryRowContext(ctx, getUser, id)
	var i User
	err := row.Scan(&i.ID, &i.Username, &i.Password)
	return i, err
}

const getUserName = `-- name: GetUserName :one
SELECT username FROM users
WHERE id = $1 LIMIT 1
`

func (q *Queries) GetUserName(ctx context.Context, id int64) (string, error) {
	row := q.db.QueryRowContext(ctx, getUserName, id)
	var username string
	err := row.Scan(&username)
	return username, err
}
