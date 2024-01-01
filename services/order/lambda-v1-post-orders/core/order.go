package core

import (
	"context"
	"fmt"
	"root/services/order/lambda-shared/incoming"
	shared "root/services/order/lambda-shared/outgoing"
	"root/services/order/lambda-v1-post-orders/outgoing"
	"time"
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

func (service *OrderService) PlaceOrder(ctx context.Context, orderRequest incoming.OrderRequest) (
	incoming.OrderResponse,
	error,
) {
	creationDate := time.Now()
	orderId := NewOrderId(
		service.region,
	)

	orderItems := make([]shared.OrderItemEntity, 0)
	for _, item := range orderRequest.Items {
		orderItems = append(
			orderItems, shared.OrderItemEntity{
				OrderItemId:  0,
				OrderId:      string(orderId),
				ItemName:     item.Name,
				CreationDate: creationDate,
			},
		)
	}

	order := shared.OrderEntity{
		OrderId:       string(orderId),
		CustomerId:    orderRequest.CustomerId,
		OrderWorkflow: "default_workflow",
		CreationDate:  creationDate,
		OrderStatus:   string(incoming.OrderPlaced),
	}

	err := service.orderRepository.SaveOrder(ctx, order, orderItems)
	if err != nil {
		return incoming.OrderResponse{}, fmt.Errorf("failed to save order: %w", err)
	}

	orderItemResponses := make([]incoming.OrderItemResponse, 0)
	for _, orderItem := range orderItems {
		orderItemResponses = append(
			orderItemResponses, incoming.OrderItemResponse{
				Name: orderItem.ItemName,
			},
		)
	}

	return incoming.OrderResponse{
		OrderId:      string(orderId),
		CustomerId:   orderRequest.CustomerId,
		CreationDate: creationDate,
		Status:       incoming.OrderPlaced,
		Items:        orderItemResponses,
	}, nil
}
