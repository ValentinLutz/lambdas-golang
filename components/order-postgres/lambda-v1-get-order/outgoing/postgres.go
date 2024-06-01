package outgoing

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"root/components/order-postgres/lambda-v1-get-order/outgoing/postgres"
	"time"
)

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate --file ../sqlc.yaml

var (
	ErrOrderNotFound = errors.New("order not found")
)

type OrderRepository struct {
	*postgres.Queries
}

func NewOrderRepository(db postgres.DBTX) *OrderRepository {
	queries := postgres.New(db)

	return &OrderRepository{
		Queries: queries,
	}
}

type Order struct {
	OrderId     string
	CustomerId  uuid.UUID
	CreatedAt   time.Time
	OrderStatus string
}

type OrderItem struct {
	ItemName string
}

func (orderRepo OrderRepository) GetOrderByOrderId(ctx context.Context, orderId string) (Order, error) {
	result, err := orderRepo.FindOrderByOrderId(ctx, orderId)
	if errors.Is(err, pgx.ErrNoRows) {
		return Order{}, fmt.Errorf("%w: order id %s: %w", ErrOrderNotFound, orderId, err)
	}
	if err != nil {
		return Order{}, fmt.Errorf("failed to fetch order from database: %w", err)
	}

	return Order{
		OrderId:     result.OrderID,
		CustomerId:  result.CustomerID.Bytes,
		CreatedAt:   result.CreatedAt.Time.UTC(),
		OrderStatus: result.Status,
	}, nil
}

func (orderRepo OrderRepository) GetOrderItemsByOrderId(ctx context.Context, orderId string) ([]OrderItem, error) {
	results, err := orderRepo.FindOrderItemsByOrderId(ctx, orderId)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("%w: order id %d: %w", ErrOrderNotFound, orderId, err)
	}
	if err != nil {
		return nil, fmt.Errorf("failed to fetch order items from database: %w", err)
	}

	var orderItems []OrderItem
	for _, orderItem := range results {
		orderItems = append(orderItems, OrderItem{
			ItemName: orderItem.Name,
		})
	}

	return orderItems, nil
}
