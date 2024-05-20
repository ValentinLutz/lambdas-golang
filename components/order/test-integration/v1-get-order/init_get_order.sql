INSERT INTO order_service.order
    (order_id, customer_id, order_workflow, creation_date, order_status)
VALUES ('01HK5W7JF32CM-EU-0GVJSF5RFM7PN', '44bd6239-7e3d-4d4a-90a0-7d4676a00f5c', 'default_workflow',
        '1970-01-01 00:00:00 +00:00', 'order_placed');

INSERT INTO order_service.order_item
    (order_id, creation_date, order_item_name)
VALUES ('01HK5W7JF32CM-EU-0GVJSF5RFM7PN', '1970-01-01 00:00:00 +00:00', 'orange'),
       ('01HK5W7JF32CM-EU-0GVJSF5RFM7PN', '1970-01-01 00:00:00 +00:00', 'banana'),
       ('01HK5W7JF32CM-EU-0GVJSF5RFM7PN', '1970-01-01 00:00:00 +00:00', 'apple');
