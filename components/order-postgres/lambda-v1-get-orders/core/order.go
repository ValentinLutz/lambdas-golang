package core

import (
	"context"
	"errors"
	"fmt"
	"root/components/order-postgres/lambda-v1-get-orders/outgoing"
	"root/components/order-postgres/lambda-v1-get-orders/outgoing/postgres"

	"github.com/google/uuid"
)

var (
	ErrInvalidOffset = errors.New("offset must be greater than or equal to 0")
	ErrInvalidLimit  = errors.New("limit must be greater than 0")
)

type OrderService struct {
	orderRepository *outgoing.OrderRepository
}

func NewOrderService(orderRepository *outgoing.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (service *OrderService) GetOrders(ctx context.Context, offset int, limit int, customerId *uuid.UUID) ([]postgres.OrderServiceOrder, []postgres.OrderServiceOrderItem, error) {
	if offset < 0 {
		return nil, nil, ErrInvalidOffset
	}
	if limit <= 0 {
		return nil, nil, ErrInvalidLimit
	}

	if customerId != nil {
		orderEntities, orderItemEntities, err := service.orderRepository.GetOrdersByCustomerId(ctx, *customerId, offset, limit)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get orders by customer id: %w", err)
		}

		return orderEntities, orderItemEntities, nil
	}

	orderEntities, orderItemEntities, err := service.orderRepository.GetOrders(ctx, offset, limit)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get orders: %w", err)
	}

	return orderEntities, orderItemEntities, nil
}
