package outgoing

import (
	"context"
	"root/services/order/lambda-shared/outgoing"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type OrderRepository struct {
	*sqlx.DB
}

func NewOrderRepository(database *sqlx.DB) *OrderRepository {
	return &OrderRepository{DB: database}
}

func (orderRepository *OrderRepository) FindAllOrdersByCustomerId(
	ctx context.Context,
	customerId uuid.UUID,
	offset int,
	limit int,
) ([]outgoing.OrderEntity, []outgoing.OrderItemEntity, error) {
	var orderEntities []outgoing.OrderEntity
	err := orderRepository.SelectContext(
		ctx,
		&orderEntities,
		"SELECT order_id, customer_id, creation_date, order_status FROM order_service.order WHERE customer_id = $1 ORDER BY creation_date OFFSET $2 LIMIT $3",
		customerId,
		offset,
		limit,
	)
	if err != nil {
		return nil, nil, err
	}

	var orderIds []string
	for _, order := range orderEntities {
		orderIds = append(orderIds, order.OrderId)
	}

	var orderItemEntities []outgoing.OrderItemEntity
	err = orderRepository.SelectContext(
		ctx,
		&orderItemEntities,
		"SELECT order_item_id, order_id, creation_date, order_item_name FROM order_service.order_item WHERE order_id = ANY($1)",
		pq.Array(orderIds),
	)
	if err != nil {
		return nil, nil, err
	}

	return orderEntities, orderItemEntities, nil
}

func (orderRepository *OrderRepository) FindAllOrders(ctx context.Context, offset int, limit int) (
	[]outgoing.OrderEntity,
	[]outgoing.OrderItemEntity,
	error,
) {
	var orderEntities []outgoing.OrderEntity
	err := orderRepository.SelectContext(
		ctx,
		&orderEntities,
		"SELECT order_id, customer_id, creation_date, order_status FROM order_service.order ORDER BY creation_date OFFSET $1 LIMIT $2",
		offset, limit,
	)
	if err != nil {
		return nil, nil, err
	}

	var orderIds []string
	for _, order := range orderEntities {
		orderIds = append(orderIds, order.OrderId)
	}

	var orderItemEntities []outgoing.OrderItemEntity
	err = orderRepository.SelectContext(
		ctx,
		&orderItemEntities,
		"SELECT order_item_id, order_id, creation_date, order_item_name FROM order_service.order_item WHERE order_id = ANY($1)",
		pq.Array(orderIds),
	)
	if err != nil {
		return nil, nil, err
	}

	return orderEntities, orderItemEntities, nil
}
