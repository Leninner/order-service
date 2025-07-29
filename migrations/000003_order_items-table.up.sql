create table order_items (
    id varchar(50) primary key,
    order_id uuid not null references orders(id) on delete cascade,
    product_id uuid not null,
    quantity smallint not null,
    price decimal(10,2) not null,
    sub_total decimal(10,2) not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);

create index idx_order_items_order_id on order_items(order_id); 