// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.13.0

package db

import (
	"context"
)

type Querier interface {
	AddAuth(ctx context.Context, arg AddAuthParams) error
	GetAuth(ctx context.Context, username string) (Auth, error)
}

var _ Querier = (*Queries)(nil)
