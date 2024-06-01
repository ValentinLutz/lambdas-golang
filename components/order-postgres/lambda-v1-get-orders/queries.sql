-- name: FindOrdersByCustomerId :many
SELECT *
FROM order_service.order
WHERE customer_id = $1
ORDER BY created_at
OFFSET $2 LIMIT $3;

-- name: FindAllOrders :many
SELECT *
FROM order_service.order
ORDER BY created_at
OFFSET $1 LIMIT $2;

-- name: FindOrderItemsByOrderIds :many
SELECT *
FROM order_service.order_item
WHERE order_id = ANY (@order_ids::varchar[]);