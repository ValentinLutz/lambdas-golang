CREATE TABLE IF NOT EXISTS order_service.order
(
    order_id    VARCHAR     NOT NULL UNIQUE,
    customer_id UUID        NOT NULL,
    workflow    VARCHAR     NOT NULL,
    status      VARCHAR     NOT NULL,
    created_at  TIMESTAMPTZ NOT NULL,
    modified_at TIMESTAMPTZ NOT NULL DEFAULT now(),

    PRIMARY KEY (order_id)
);

CREATE INDEX order_id_idx
    ON order_service.order (order_id);

CREATE INDEX order_created_at_idx
    ON order_service.order (created_at);

CREATE TRIGGER update_order_modified_at
    BEFORE UPDATE
    ON order_service.order
    FOR EACH ROW
EXECUTE PROCEDURE order_service.update_modified_at();
