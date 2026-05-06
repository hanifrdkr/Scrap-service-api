-- auto-generated definition
create table users
(
    id         uuid                     default gen_random_uuid() not null
        primary key,
    email      varchar(320)                                       not null,
    name       varchar(225),
    created_at timestamp with time zone default CURRENT_TIMESTAMP,
    updated_at timestamp with time zone default CURRENT_TIMESTAMP,
    deleted_at timestamp with time zone
);

insert into users(email, name, created_at, updated_at) values ('admin@example.com', 'admin', current_timestamp, current_timestamp);