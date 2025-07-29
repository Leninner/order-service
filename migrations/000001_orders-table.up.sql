create table orders (
    id uuid primary key default gen_random_uuid(),
    customer_id uuid not null,
    restaurant_id uuid not null,
    tracking_id uuid not null,
    price decimal(10,2) not null,
    order_status varchar(50) not null,
    failure_messages jsonb default '[]'::jsonb,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
);