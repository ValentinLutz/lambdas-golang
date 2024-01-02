package outgoing

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	shared "root/services/order/lambda-shared/outgoing"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	*sqlx.DB
}

var (
	ErrOrderNotFound = fmt.Errorf("order not found")
)

func NewOrderRepository(database *sqlx.DB) *OrderRepository {
	return &OrderRepository{DB: database}
}

func (orderRepository *OrderRepository) FindOrderByOrderId(ctx context.Context, orderId string) (
	shared.OrderEntity,
	[]shared.OrderItemEntity,
	error,
) {
	var orderEntity shared.OrderEntity
	err := orderRepository.GetContext(
		ctx,
		&orderEntity,
		`SELECT order_id, customer_id, creation_date, order_status 
				FROM order_service.order 
				WHERE order_id = $1`,
		orderId,
	)
	if errors.Is(err, sql.ErrNoRows) {
		return shared.OrderEntity{}, nil, fmt.Errorf("%w: order id %s: %w", ErrOrderNotFound, orderId, err)
	}
	if err != nil {
		return shared.OrderEntity{}, nil, fmt.Errorf("failed to fetch order from database: %w", err)
	}

	var orderItemEntities []shared.OrderItemEntity
	err = orderRepository.SelectContext(
		ctx,
		&orderItemEntities,
		`SELECT order_item_id, order_id, creation_date, order_item_name 
				FROM order_service.order_item 
				WHERE order_id = $1`,
		orderId,
	)
	if err != nil {
		return shared.OrderEntity{}, nil, fmt.Errorf("failed to fetch order items from database: %w", err)
	}

	return orderEntity, orderItemEntities, nil
}
