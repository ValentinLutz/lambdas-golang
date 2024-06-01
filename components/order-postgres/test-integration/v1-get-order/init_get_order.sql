INSERT INTO order_service.order
    (order_id, customer_id, workflow, created_at, status)
VALUES ('01HK5W7JF32CM-EU-0GVJSF5RFM7PN', '44bd6239-7e3d-4d4a-90a0-7d4676a00f5c', 'default_workflow',
        '1970-01-01 00:00:00 +00:00', 'order_placed');

INSERT INTO order_service.order_item
    (order_item_id, order_id, created_at, name)
VALUES ('01HZA1GWPT9QFGWCM54M4JJ4A8', '01HK5W7JF32CM-EU-0GVJSF5RFM7PN', '1970-01-01 00:00:00 +00:00', 'orange'),
       ('01HZA1HH1CZJKPD0HZ5PE52H9X', '01HK5W7JF32CM-EU-0GVJSF5RFM7PN', '1970-01-01 00:00:00 +00:00', 'banana'),
       ('01HZA1HPJ6737E9PQ0Y60GPAE7', '01HK5W7JF32CM-EU-0GVJSF5RFM7PN', '1970-01-01 00:00:00 +00:00', 'apple');
