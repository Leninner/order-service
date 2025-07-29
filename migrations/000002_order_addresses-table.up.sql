create table order_addresses (
    id uuid primary key default gen_random_uuid(),
    order_id uuid not null unique references orders(id) on delete cascade,
    street varchar(255) not null,
    postal_code varchar(20) not null,
    city varchar(100) not null,
    created_at timestamp with time zone default now(),
    updated_at timestamp with time zone default now()
); 