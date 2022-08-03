// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"

	"github.com/google/uuid"
)

type Querier interface {
	AddUser(ctx context.Context, arg AddUserParams) error
	GetUser(ctx context.Context, id uuid.UUID) (User, error)
}

var _ Querier = (*Queries)(nil)
