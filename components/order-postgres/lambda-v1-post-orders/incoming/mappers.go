package incoming

import (
	"root/components/order-postgres/lambda-v1-post-orders/outgoing/postgres"
)

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config ../../api-definition/oapi-codgen.yaml ../../api-definition/order-api-v1.yaml

func NewOrderResponse(order postgres.OrderServiceOrder, orderItems []postgres.OrderServiceOrderItem) OrderResponse {
	orderItemResponses := make([]OrderItemResponse, 0)
	for _, orderItem := range orderItems {
		orderItemResponses = append(
			orderItemResponses, OrderItemResponse{
				Name: orderItem.Name,
			},
		)
	}

	return OrderResponse{
		OrderId:    order.OrderID,
		CustomerId: order.CustomerID.Bytes,
		Items:      orderItemResponses,
		Status:     OrderStatus(order.Status),
	}
}
