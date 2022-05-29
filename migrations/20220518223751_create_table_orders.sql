-- +goose Up
-- +goose StatementBegin
create type order_status as enum ('new', 'accepted', 'in_proccess', 'delivered');
create table orders(
    id uuid primary key,
    from_addr text not null,
    to_addr text not null,
    status order_status not null default 'new'::order_status,
    price smallint not null,
    sender_name text not null,
    sender_passport_num text not null,
    recipient_name text not null,
    weight smallint not null,
    created_at timestamptz not null default now(),
    created_by uuid not null
);
create index orders_sender_passport_num_idx on orders using btree(sender_passport_num);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table orders;
-- +goose StatementEnd