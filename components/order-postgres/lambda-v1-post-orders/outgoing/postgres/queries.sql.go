// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: queries.sql

package postgres

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const saveOrder = `-- name: SaveOrder :exec
INSERT INTO order_service.order
    (order_id, customer_id, created_at, status, workflow)
VALUES ($1, $2, $3, $4, $5)
`

type SaveOrderParams struct {
	OrderID    string
	CustomerID pgtype.UUID
	CreatedAt  pgtype.Timestamptz
	Status     string
	Workflow   string
}

func (q *Queries) SaveOrder(ctx context.Context, arg SaveOrderParams) error {
	_, err := q.db.Exec(ctx, saveOrder,
		arg.OrderID,
		arg.CustomerID,
		arg.CreatedAt,
		arg.Status,
		arg.Workflow,
	)
	return err
}

type SaveOrderItemsParams struct {
	OrderItemID string
	OrderID     string
	CreatedAt   pgtype.Timestamptz
	Name        string
}
