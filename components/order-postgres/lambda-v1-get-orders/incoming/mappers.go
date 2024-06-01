package incoming

import (
	"root/components/order-postgres/lambda-v1-get-orders/outgoing/postgres"
)

//go:generate go run github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen --config ../../api-definition/oapi-codgen.yaml ../../api-definition/order-api-v1.yaml

func NewOrdersResponse(orders []postgres.OrderServiceOrder, orderItems []postgres.OrderServiceOrderItem) OrdersResponse {
	orderIdToOrderItems := make(map[string][]OrderItemResponse)
	for _, orderItemEntity := range orderItems {
		orderIdToOrderItems[orderItemEntity.OrderID] = append(
			orderIdToOrderItems[orderItemEntity.OrderID], OrderItemResponse{
				Name: orderItemEntity.Name,
			},
		)
	}

	ordersResponse := make([]OrderResponse, 0)
	for _, order := range orders {
		ordersResponse = append(
			ordersResponse, OrderResponse{
				OrderId:    order.OrderID,
				CustomerId: order.CustomerID.Bytes,
				CreatedAt:  order.CreatedAt.Time.UTC(),
				Status:     OrderStatus(order.Status),
				Items:      orderIdToOrderItems[order.OrderID],
			},
		)
	}

	return ordersResponse
}
