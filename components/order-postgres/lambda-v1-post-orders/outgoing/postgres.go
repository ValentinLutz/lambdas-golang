package outgoing

import (
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"golang.org/x/net/context"
	"root/components/order-postgres/lambda-v1-post-orders/outgoing/postgres"
)

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate --file ../sqlc.yaml

type OrderRepository struct {
	*postgres.Queries
	*pgxpool.Pool
}

func NewOrderRepository(db *pgxpool.Pool) *OrderRepository {
	queries := postgres.New(db)

	return &OrderRepository{
		Queries: queries,
		Pool:    db,
	}
}

func (orderRepo OrderRepository) SaveOrder(ctx context.Context, order postgres.OrderServiceOrder, orderItems []postgres.OrderServiceOrderItem) error {
	tx, err := orderRepo.Pool.Begin(ctx)
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}
	defer tx.Rollback(ctx)

	queries := orderRepo.WithTx(tx)

	err = queries.SaveOrder(ctx, postgres.SaveOrderParams{
		OrderID:    order.OrderID,
		CustomerID: order.CustomerID,
		CreatedAt:  order.CreatedAt,
		Status:     order.Status,
		Workflow:   order.Workflow,
	})
	if err != nil {
		return fmt.Errorf("failed to save order: %w", err)
	}

	var orderItemsParams []postgres.SaveOrderItemsParams
	for _, orderItem := range orderItems {
		orderItemsParams = append(orderItemsParams, postgres.SaveOrderItemsParams{
			OrderItemID: orderItem.OrderItemID,
			OrderID:     orderItem.OrderID,
			CreatedAt:   orderItem.CreatedAt,
			Name:        orderItem.Name,
		})
	}
	_, err = queries.SaveOrderItems(ctx, orderItemsParams)
	if err != nil {
		return fmt.Errorf("failed to save order items: %w", err)
	}

	return tx.Commit(ctx)
}
