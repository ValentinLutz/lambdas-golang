package outgoing

import (
	"context"
	"errors"
	"root/services/order/lambda-shared/outgoing"

	"github.com/jmoiron/sqlx"
)

type OrderRepository struct {
	*sqlx.DB
}

func NewOrderRepository(database *sqlx.DB) *OrderRepository {
	return &OrderRepository{DB: database}
}

func (orderRepository *OrderRepository) SaveOrder(
	ctx context.Context,
	orderEntity outgoing.OrderEntity,
	orderItemEntities []outgoing.OrderItemEntity,
) error {
	return orderRepository.execTx(
		ctx, func(tx *sqlx.Tx) error {
			_, err := tx.NamedExec(
				`INSERT INTO order_service.order (order_id, customer_id, creation_date, order_status, order_workflow) 
							VALUES (:order_id, :customer_id, :creation_date, :order_status, :order_workflow);`,
				orderEntity,
			)
			if err != nil {
				return err
			}

			_, err = tx.NamedExec(
				`INSERT INTO order_service.order_item (order_id, creation_date, order_item_name) 
							VALUES (:order_id, :creation_date, :order_item_name);`,
				orderItemEntities,
			)
			if err != nil {
				return err
			}

			return nil
		},
	)
}

func (orderRepository *OrderRepository) execTx(ctx context.Context, fn func(*sqlx.Tx) error) error {
	tx, err := orderRepository.BeginTxx(ctx, nil)
	if err != nil {
		return err
	}

	err = fn(tx)
	if err != nil {
		return errors.Join(err, tx.Rollback())
	}

	return tx.Commit()
}
