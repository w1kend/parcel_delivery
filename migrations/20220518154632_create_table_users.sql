-- +goose Up
-- +goose StatementBegin
create table users(
    id uuid primary key,
    name text not null,
    email text not null,
    password text not null,
    created_at timestamptz not null default now()
);
create index users_email_idx on users using btree(email);
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd