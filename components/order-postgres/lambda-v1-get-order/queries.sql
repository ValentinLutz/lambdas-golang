-- name: FindOrderByOrderId :one
SELECT *
FROM order_service.order
WHERE order_id = $1;


-- name: FindOrderItemsByOrderId :many
SELECT *
FROM order_service.order_item
WHERE order_id = $1;
