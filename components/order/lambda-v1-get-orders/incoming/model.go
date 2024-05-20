package incoming

import "root/components/order/lambda-v1-get-orders/outgoing"

//go:generate oapi-codegen --config ../../api-definition/oapi-codgen.yaml ../../api-definition/order-api-v1.yaml

func NewOrdersResponse(orders []outgoing.OrderEntity, orderItems []outgoing.OrderItemEntity) OrdersResponse {
	orderIdToOrderItems := make(map[string][]OrderItemResponse)
	for _, orderItemEntity := range orderItems {
		orderIdToOrderItems[orderItemEntity.OrderId] = append(
			orderIdToOrderItems[orderItemEntity.OrderId], OrderItemResponse{
				Name: orderItemEntity.ItemName,
			},
		)
	}

	ordersResponse := make([]OrderResponse, 0)
	for _, order := range orders {
		ordersResponse = append(
			ordersResponse, OrderResponse{
				OrderId:      order.OrderId,
				CustomerId:   order.CustomerId,
				CreationDate: order.CreationDate,
				Status:       OrderStatus(order.OrderStatus),
				Items:        orderIdToOrderItems[order.OrderId],
			},
		)
	}

	return ordersResponse
}
