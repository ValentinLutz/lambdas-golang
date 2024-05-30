package incoming

import "root/components/order/lambda-v1-get-order/outgoing"

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config ../../api-definition/oapi-codgen.yaml ../../api-definition/order-api-v1.yaml

func NewOrderResponse(order outgoing.OrderEntity, orderItems []outgoing.OrderItemEntity) OrderResponse {
	orderItemResponses := make([]OrderItemResponse, 0)
	for _, orderItem := range orderItems {
		orderItemResponses = append(
			orderItemResponses,
			OrderItemResponse{
				Name: orderItem.ItemName,
			},
		)
	}

	orderResponse := OrderResponse{
		CreationDate: order.CreationDate,
		CustomerId:   order.CustomerId,
		Items:        orderItemResponses,
		OrderId:      order.OrderId,
		Status:       OrderStatus(order.OrderStatus),
	}

	return orderResponse
}
