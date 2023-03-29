create table comments(
    id serial primary key, 
    user_id integer, 
    product_id integer, 
    description text,
    created_at timestamp default current_timestamp,
    deleted_at timestamp
);