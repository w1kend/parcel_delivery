-- +goose Up
-- +goose StatementBegin
create table orders(
    id uuid primary key,
    from_addr text not null,
    to_addr text not null,
    sender_name text not null,
    sender_passport_num text not null,
    recipient_name text not null,
    weight smallint not null
);
create index orders_sender_passport_num_inx on orders using btree(sender_passport_num);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table orders;
-- +goose StatementEnd