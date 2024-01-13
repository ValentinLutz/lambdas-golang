package core

import (
	"context"
	"fmt"
	"root/services/order/lambda-v1-get-order/outgoing"
)

type OrderService struct {
	orderRepository *outgoing.OrderRepository
}

func NewOrderService(orderRepository *outgoing.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (service *OrderService) GetOrder(ctx context.Context, orderId string) (outgoing.OrderEntity, []outgoing.OrderItemEntity, error) {
	order, orderItems, err := service.orderRepository.FindOrderByOrderId(ctx, orderId)
	if err != nil {
		return outgoing.OrderEntity{}, nil, fmt.Errorf("failed to get order: %w", err)
	}

	return order, orderItems, nil
}
