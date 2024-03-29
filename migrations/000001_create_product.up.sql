create table if not exists products_exam(
    id text primary key,
    owner_id text ,
    name text not null,
    description text not null,
    price integer not null,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);