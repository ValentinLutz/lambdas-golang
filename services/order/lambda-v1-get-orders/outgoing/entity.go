package outgoing

import (
	"time"

	"github.com/google/uuid"
)

type OrderEntity struct {
	OrderId       string    `db:"order_id"`
	CustomerId    uuid.UUID `db:"customer_id"`
	CreationDate  time.Time `db:"creation_date"`
	OrderStatus   string    `db:"order_status"`
	OrderWorkflow string    `db:"order_workflow"`
}

type OrderItemEntity struct {
	OrderItemId  int       `db:"order_item_id"`
	OrderId      string    `db:"order_id"`
	ItemName     string    `db:"order_item_name"`
	CreationDate time.Time `db:"creation_date"`
}
