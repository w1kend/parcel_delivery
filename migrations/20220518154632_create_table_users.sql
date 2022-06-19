-- +goose Up
-- +goose StatementBegin
create type user_role as enum('admin', 'client', 'courier');
create table users(
    id uuid primary key,
    name text not null,
    email text not null,
    password text not null,
    role user_role not null,
    created_at timestamptz not null default now()
);
create index users_email_idx on users using btree(email);
insert into users(id, name, email, password, role)
values (
        gen_random_uuid(),
        'admin',
        'admin@app.com',
        '$2a$12$Gay6BGVt5wzeWQ3CxDxXjuA0nrNqP.PoFj4QmOl/pj96T4.h5gx02',
        'admin'::user_role
    );
-- +goose StatementEnd
-- +goose Down
-- +goose StatementBegin
drop table users;
-- +goose StatementEnd