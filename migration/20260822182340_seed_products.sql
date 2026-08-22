-- +goose Up

INSERT INTO product_specs (guid, name, "desc")
VALUES
    ('c8cd15fc-8ac4-4593-b83c-ff48e7a04866', 'disc brake', 'brake for 23 inch wheel'),
    ('4369bbde-810e-4ec3-9a35-012a209c4214', 'brambo disc pad front', 'pad 10 inch'),
    ('d95cefcb-33cf-443a-ba07-ce8b3feda1ae', 'brambo disc pad rear', 'pad 15 inch');

INSERT INTO orders (guid, status_id)
VALUES
    ('2080dce7-0e1b-4911-ad54-7ae535fd3bdc', 1),
    ('37a419e6-dd82-418a-906f-7d26a2d2664b', 2);

INSERT INTO products (guid, article_id, order_guid, spec_guid)
VALUES
    ('ac363c05-85cb-48aa-8f5a-7a078e8255d0', 1165547, '2080dce7-0e1b-4911-ad54-7ae535fd3bdc', 'c8cd15fc-8ac4-4593-b83c-ff48e7a04866'),
    ('c9e03a18-2551-40b1-a6ee-d5c7402e2c56', 1157673, '2080dce7-0e1b-4911-ad54-7ae535fd3bdc', '4369bbde-810e-4ec3-9a35-012a209c4214'),
    ('8479d2b2-5558-4149-ae35-4eba399c2a85', 1168746, '37a419e6-dd82-418a-906f-7d26a2d2664b', 'd95cefcb-33cf-443a-ba07-ce8b3feda1ae');
