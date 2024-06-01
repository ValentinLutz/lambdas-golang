package outgoing

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"golang.org/x/net/context"
	"root/components/order-postgres/lambda-v1-get-orders/outgoing/postgres"
)

//go:generate go run github.com/sqlc-dev/sqlc/cmd/sqlc generate --file ../sqlc.yaml

type OrderRepository struct {
	*postgres.Queries
}

func NewOrderRepository(db postgres.DBTX) *OrderRepository {
	queries := postgres.New(db)

	return &OrderRepository{
		Queries: queries,
	}
}

func (orderRepo OrderRepository) GetOrdersByCustomerId(ctx context.Context, customerId uuid.UUID, offset int, limit int) ([]postgres.OrderServiceOrder, []postgres.OrderServiceOrderItem, error) {
	orders, err := orderRepo.Queries.FindOrdersByCustomerId(ctx, postgres.FindOrdersByCustomerIdParams{
		CustomerID: pgtype.UUID{
			Bytes: customerId,
			Valid: true,
		},
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch orders from database: %w", err)
	}

	var orderIds []string
	for _, order := range orders {
		orderIds = append(orderIds, order.OrderID)
	}

	orderItems, err := orderRepo.Queries.FindOrderItemsByOrderIds(ctx, orderIds)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch order items from database: %w", err)
	}

	return orders, orderItems, nil
}

func (orderRepo OrderRepository) GetOrders(ctx context.Context, offset int, limit int) ([]postgres.OrderServiceOrder, []postgres.OrderServiceOrderItem, error) {
	orders, err := orderRepo.FindAllOrders(ctx, postgres.FindAllOrdersParams{
		Offset: int32(offset),
		Limit:  int32(limit),
	})
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch orders from database: %w", err)
	}

	var orderIds []string
	for _, order := range orders {
		orderIds = append(orderIds, order.OrderID)
	}

	orderItems, err := orderRepo.FindOrderItemsByOrderIds(ctx, orderIds)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to fetch order items from database: %w", err)
	}

	return orders, orderItems, nil
}
