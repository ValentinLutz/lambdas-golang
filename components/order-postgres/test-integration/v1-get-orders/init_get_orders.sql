INSERT INTO order_service.order
    (order_id, customer_id, workflow, created_at, status)
VALUES ('01HK5VD1NN2V0-EU-EHHDDPYPXDDRA', '44bd6239-7e3d-4d4a-90a0-7d4676a00f5c', 'default_workflow',
        '1970-01-01 00:00:00 +00:00', 'order_placed'),

       ('01HK5VZG8ZY60-EU-VNB5J6P15EPB0', '44bd6239-7e3d-4d4a-90a0-7d4676a00f5c', 'default_workflow',
        '1970-01-01 00:00:00 +00:00', 'order_in_progress'),

       ('01HK5VZG90SWD-EU-MK2PN9ZV2FRN8', '44bd6239-7e3d-4d4a-90a0-7d4676a00f5c', 'default_workflow',
        '1970-01-01 00:00:00 +00:00', 'order_canceled'),

       ('01HK5VZG93FC9-EU-QYR4T29PR68M3', 'cd8641e3-196a-4045-a0b6-706a80e48262', 'default_workflow',
        '1970-01-01 00:00:00 +00:00', 'order_completed');

INSERT INTO order_service.order_item
    (order_item_id, order_id, created_at, name)
VALUES ('01HZADD8ZY945HJ1GC4X6XZ8YT', '01HK5VD1NN2V0-EU-EHHDDPYPXDDRA', '1970-01-01 00:00:00 +00:00', 'orange'),
       ('01HZADDH6DAJHEWM1RE3XQZYAT', '01HK5VD1NN2V0-EU-EHHDDPYPXDDRA', '1970-01-01 00:00:00 +00:00', 'banana'),

       ('01HZADDRBGCVQWR958AMPVPZBX', '01HK5VZG8ZY60-EU-VNB5J6P15EPB0', '1970-01-01 00:00:00 +00:00', 'chocolate'),

       ('01HZADDZAF5D5A1S6TQB7MJV3V', '01HK5VZG90SWD-EU-MK2PN9ZV2FRN8', '1970-01-01 00:00:00 +00:00', 'marshmallow'),

       ('01HZADE77SVMSXJDZS3VEZY5HX', '01HK5VZG93FC9-EU-QYR4T29PR68M3', '1970-01-01 00:00:00 +00:00', 'apple');