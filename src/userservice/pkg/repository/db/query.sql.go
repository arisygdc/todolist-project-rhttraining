// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0
// source: query.sql

package db

import (
	"context"

	"github.com/google/uuid"
)

const addUser = `-- name: AddUser :exec
INSERT INTO users (id, first_name, last_name, auth_id) VALUES ($1, $2, $3, $4)
`

type AddUserParams struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	AuthID    uuid.UUID `json:"auth_id"`
}

func (q *Queries) AddUser(ctx context.Context, arg AddUserParams) error {
	_, err := q.db.Exec(ctx, addUser,
		arg.ID,
		arg.FirstName,
		arg.LastName,
		arg.AuthID,
	)
	return err
}

const getUser = `-- name: GetUser :one
SELECT id, first_name, last_name, auth_id FROM users WHERE auth_id = $1
`

func (q *Queries) GetUser(ctx context.Context, authID uuid.UUID) (User, error) {
	row := q.db.QueryRow(ctx, getUser, authID)
	var i User
	err := row.Scan(
		&i.ID,
		&i.FirstName,
		&i.LastName,
		&i.AuthID,
	)
	return i, err
}
