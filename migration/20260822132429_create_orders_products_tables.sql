-- +goose Up

CREATE TABLE product_specs (
    guid UUID PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    "desc" VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE orders (
    guid UUID PRIMARY KEY,
    status_id SMALLINT NOT NULL CHECK (status_id >= 0),
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW()
);

CREATE TABLE products (
    guid UUID PRIMARY KEY,
    article_id BIGINT NOT NULL CHECK (article_id >= 0),
    order_guid UUID NOT NULL,
    spec_guid UUID NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    CONSTRAINT fk_product_specs
        FOREIGN KEY (spec_guid)
        REFERENCES product_specs(guid),
    CONSTRAINT fk_products_order
        FOREIGN KEY (order_guid)
        REFERENCES orders(guid)
);


-- Indexes

CREATE INDEX idx_product_specs_name
    ON product_specs(name);

CREATE INDEX idx_product_specs_created_at
    ON product_specs(created_at);

CREATE INDEX idx_products_article_id
    ON products(article_id);

CREATE INDEX idx_products_created_at
    ON products(created_at);

CREATE INDEX idx_orders_status
    ON orders(status_id);

CREATE INDEX idx_orders_created_at
    ON orders(created_at);

-- +goose Down

DROP TABLE orders;
DROP TABLE products;
DROP TABLE product_spec;