package core

import (
	"context"
	"fmt"
	"root/services/order/lambda-shared/incoming"
	shared "root/services/order/lambda-shared/outgoing"
	"root/services/order/lambda-v1-get-orders/outgoing"

	"github.com/google/uuid"
)

var (
	ErrInvalidOffset = fmt.Errorf("offset must be greater than or equal to 0")
	ErrInvalidLimit  = fmt.Errorf("limit must be greater than 0")
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
	incoming.OrdersResponse,
	error,
) {
	if offset < 0 {
		return nil, ErrInvalidOffset
	}
	if limit <= 0 {
		return nil, ErrInvalidLimit
	}

	if customerId != nil {
		orderEntities, orderItemEntities, err := service.orderRepository.FindAllOrdersByCustomerId(ctx, *customerId, offset, limit)
		if err != nil {
			return nil, fmt.Errorf("failed to get orders by customer id: %w", err)
		}

		return newOrdersResponse(orderEntities, orderItemEntities), nil
	}

	orderEntities, orderItemEntities, err := service.orderRepository.FindAllOrders(ctx, offset, limit)
	if err != nil {
		return nil, fmt.Errorf("failed to get orders: %w", err)
	}

	return newOrdersResponse(orderEntities, orderItemEntities), nil
}

func newOrdersResponse(
	orderEntities []shared.OrderEntity,
	orderItemEntities []shared.OrderItemEntity,
) incoming.OrdersResponse {
	orderIdToOrderItems := make(map[string][]incoming.OrderItemResponse)
	for _, orderItemEntity := range orderItemEntities {
		orderIdToOrderItems[orderItemEntity.OrderId] = append(
			orderIdToOrderItems[orderItemEntity.OrderId], incoming.OrderItemResponse{
				Name: orderItemEntity.ItemName,
			},
		)
	}

	orders := make([]incoming.OrderResponse, 0)
	for _, orderEntity := range orderEntities {
		orders = append(
			orders, incoming.OrderResponse{
				OrderId:      orderEntity.OrderId,
				CustomerId:   orderEntity.CustomerId,
				CreationDate: orderEntity.CreationDate,
				Status:       incoming.OrderStatus(orderEntity.OrderStatus),
				Items:        orderIdToOrderItems[orderEntity.OrderId],
			},
		)
	}

	return orders
}
