package core

import (
	"context"
	"root/components/order-postgres/lambda-v1-get-order/outgoing"
)

type OrderService struct {
	orderRepository *outgoing.OrderRepository
}

func NewOrderService(orderRepository *outgoing.OrderRepository) *OrderService {
	return &OrderService{
		orderRepository: orderRepository,
	}
}

func (service *OrderService) GetOrder(ctx context.Context, orderId string) (outgoing.Order, []outgoing.OrderItem, error) {
	order, err := service.orderRepository.GetOrderByOrderId(ctx, orderId)
	if err != nil {
		return outgoing.Order{}, nil, err
	}

	orderItems, err := service.orderRepository.GetOrderItemsByOrderId(ctx, orderId)
	if err != nil {
		return outgoing.Order{}, nil, err
	}

	return order, orderItems, nil
}
