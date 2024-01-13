package core

import (
	"context"
	"fmt"
	"root/services/order/lambda-v1-post-orders/outgoing"
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
	outgoing.OrderEntity,
	[]outgoing.OrderItemEntity,
	error,
) {
	creationDate := time.Now()
	orderId := NewOrderId(
		service.region,
	)

	orderItems := make([]outgoing.OrderItemEntity, 0)
	for _, item := range items {
		orderItems = append(
			orderItems, outgoing.OrderItemEntity{
				OrderItemId:  0,
				OrderId:      string(orderId),
				ItemName:     item,
				CreationDate: creationDate,
			},
		)
	}

	order := outgoing.OrderEntity{
		OrderId:       string(orderId),
		CustomerId:    customerId,
		OrderWorkflow: "default_workflow",
		CreationDate:  creationDate,
		OrderStatus:   "order_placed",
	}

	err := service.orderRepository.SaveOrder(ctx, order, orderItems)
	if err != nil {
		return outgoing.OrderEntity{}, nil, fmt.Errorf("failed to save order: %w", err)
	}

	return order, orderItems, nil
}
