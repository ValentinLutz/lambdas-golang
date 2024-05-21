package core

import (
	"context"
	"errors"
	"fmt"
	"root/components/order/lambda-v1-get-orders/outgoing"

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

func (service *OrderService) GetOrders(ctx context.Context, offset int, limit int, customerId *uuid.UUID) (
	[]outgoing.OrderEntity, []outgoing.OrderItemEntity,
	error,
) {
	if offset < 0 {
		return nil, nil, ErrInvalidOffset
	}
	if limit <= 0 {
		return nil, nil, ErrInvalidLimit
	}

	if customerId != nil {
		orderEntities, orderItemEntities, err := service.orderRepository.FindAllOrdersByCustomerId(ctx, *customerId, offset, limit)
		if err != nil {
			return nil, nil, fmt.Errorf("failed to get orders by customer id: %w", err)
		}

		return orderEntities, orderItemEntities, nil
	}

	orderEntities, orderItemEntities, err := service.orderRepository.FindAllOrders(ctx, offset, limit)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get orders: %w", err)
	}

	return orderEntities, orderItemEntities, nil
}
