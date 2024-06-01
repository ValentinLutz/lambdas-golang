-- name: SaveOrder :exec
INSERT INTO order_service.order
    (order_id, customer_id, created_at, status, workflow)
VALUES (@order_id, @customer_id, @created_at, @status, @workflow);

-- name: SaveOrderItems :copyfrom
INSERT INTO order_service.order_item
    (order_item_id, order_id, created_at, name)
VALUES (@order_item_id, @order_id, @created_at, @name);