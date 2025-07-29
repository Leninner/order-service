CREATE TABLE order_restaurant_m_view (
    restaurant_id UUID NOT NULL,
    product_id UUID NOT NULL,
    restaurant_name VARCHAR(255) NOT NULL,
    restaurant_active BOOLEAN NOT NULL,
    product_name VARCHAR(255) NOT NULL,
    product_price DECIMAL(10,2) NOT NULL,
    product_available BOOLEAN NOT NULL,
    PRIMARY KEY (restaurant_id, product_id)
);
