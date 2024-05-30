package incoming

import "root/components/order/lambda-v1-post-orders/outgoing"

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config ../../api-definition/oapi-codgen.yaml ../../api-definition/order-api-v1.yaml

func NewOrderResponse(order outgoing.OrderEntity, orderItems []outgoing.OrderItemEntity) OrderResponse {
	orderItemResponses := make([]OrderItemResponse, 0)
	for _, orderItem := range orderItems {
		orderItemResponses = append(
			orderItemResponses, OrderItemResponse{
				Name: orderItem.ItemName,
			},
		)
	}

	return OrderResponse{
		OrderId:      order.OrderId,
		CustomerId:   order.CustomerId,
		CreationDate: order.CreationDate,
		Status:       OrderStatus(order.OrderStatus),
		Items:        orderItemResponses,
	}
}
