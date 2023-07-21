// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: queries.sql

package sqlc

import (
	"context"

	uuid "github.com/google/uuid"
)

const listTransactions = `-- name: ListTransactions :many
SELECT id, account_id, operation_type_id, amount, event_date FROM transaction WHERE account_id = $1 ORDER BY event_date DESC
`

func (q *Queries) ListTransactions(ctx context.Context, accountID uuid.UUID) ([]Transaction, error) {
	rows, err := q.db.Query(ctx, listTransactions, accountID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Transaction
	for rows.Next() {
		var i Transaction
		if err := rows.Scan(
			&i.ID,
			&i.AccountID,
			&i.OperationTypeID,
			&i.Amount,
			&i.EventDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
