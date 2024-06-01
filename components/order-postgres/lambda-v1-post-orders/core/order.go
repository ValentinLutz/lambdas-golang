package core

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/oklog/ulid/v2"
	"root/components/order-postgres/lambda-v1-post-orders/outgoing"
	"root/components/order-postgres/lambda-v1-post-orders/outgoing/postgres"
	"time"

	"github.com/google/uuid"
)

type OrderService struct {
	region          Region
	orderRepository *outgoing.OrderRepository
}

func NewOrderService(region Region, orderRepository *outgoing.OrderRepository) *OrderService {
	return &OrderService{
		region:          region,
		orderRepository: orderRepository,
	}
}

func (service *OrderService) PlaceOrder(ctx context.Context, customerId uuid.UUID, items []string) (
	postgres.OrderServiceOrder,
	[]postgres.OrderServiceOrderItem,
	error,
) {
	CreatedAt := time.Now()
	orderId := NewOrderId(
		service.region,
	)

	orderItems := make([]postgres.OrderServiceOrderItem, 0)
	for _, item := range items {
		orderItems = append(
			orderItems, postgres.OrderServiceOrderItem{
				OrderItemID: ulid.Make().String(),
				OrderID:     string(orderId),
				Name:        item,
				CreatedAt: pgtype.Timestamptz{
					Time:  CreatedAt,
					Valid: true,
				},
			},
		)
	}

	order := postgres.OrderServiceOrder{
		OrderID: string(orderId),
		CustomerID: pgtype.UUID{
			Bytes: customerId, Valid: true,
		},
		Workflow: "default_workflow",
		CreatedAt: pgtype.Timestamptz{
			Time:  CreatedAt,
			Valid: true,
		},
		Status: "order_placed",
	}

	err := service.orderRepository.SaveOrder(ctx, order, orderItems)
	if err != nil {
		return postgres.OrderServiceOrder{}, nil, fmt.Errorf("failed to save order: %w", err)
	}

	return order, orderItems, nil
}
