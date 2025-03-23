-- +goose Up

create table if not exists users (
    id SERIAL PRIMARY KEY,
    first_name varchar(128) NOT NULL,
    last_name varchar(128) NOT NULL,
    sex varchar(16) CHECK (sex IN ('мужской', 'женский')),
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
    );

create index users_id on users (id);

-- +goose Down
drop table if exists users;
drop index users_id;