// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package sqlc

import (
	"context"

	uuid "github.com/google/uuid"
)

type Querier interface {
	ListTransactions(ctx context.Context, accountID uuid.UUID) ([]Transaction, error)
}

var _ Querier = (*Queries)(nil)