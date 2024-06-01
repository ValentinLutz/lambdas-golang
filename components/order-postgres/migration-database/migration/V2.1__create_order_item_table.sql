CREATE TABLE IF NOT EXISTS order_service.order_item
(
    order_item_id VARCHAR     NOT NULL UNIQUE,
    order_id      VARCHAR     NOT NULL,
    name          VARCHAR     NOT NULL,
    created_at    TIMESTAMPTZ NOT NULL,
    modified_at   TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (order_item_id),
    CONSTRAINT order_id_fkey FOREIGN KEY (order_id) REFERENCES order_service.order (order_id)
);

CREATE INDEX order_item_order_id_idx
    ON order_service.order_item (order_id);

CREATE TRIGGER update_order_item_modified_at
    BEFORE UPDATE
    ON order_service.order_item
    FOR EACH ROW
EXECUTE PROCEDURE order_service.update_modified_at();
