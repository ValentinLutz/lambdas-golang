package core

import (
	"context"
	"fmt"
	"root/services/order/lambda-shared/incoming"
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

func (service *OrderService) GetOrder(ctx context.Context, orderId string) (incoming.OrderResponse, error) {
	order, orderItems, err := service.orderRepository.FindOrderByOrderId(ctx, orderId)
	if err != nil {
		return incoming.OrderResponse{}, fmt.Errorf("failed to get order: %w", err)
	}

	orderItemResponses := make([]incoming.OrderItemResponse, 0)
	for _, orderItem := range orderItems {
		orderItemResponses = append(
			orderItemResponses,
			incoming.OrderItemResponse{
				Name: orderItem.ItemName,
			},
		)
	}

	orderResponse := incoming.OrderResponse{
		CreationDate: order.CreationDate,
		CustomerId:   order.CustomerId,
		Items:        orderItemResponses,
		OrderId:      order.OrderId,
		Status:       incoming.OrderStatus(order.OrderStatus),
	}

	return orderResponse, nil
}
