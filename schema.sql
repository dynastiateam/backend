create table if not exists "user"
(
    id           serial not null
        constraint user_pk
            unique,
    apartment    integer,
    email        varchar,
    password     varchar,
    phone        varchar,
    first_name   varchar,
    last_name    varchar,
    role         varchar,
    residence_id integer,
    building_id  integer
        constraint user_buildings_id_fk
            references buildings (id)
            on delete cascade
);

create index if not exists user_email_index
    on "user" (email);


create table if not exists buildings
(
    id      serial
        constraint buildings_pk
            unique,
    name    varchar,
    address varchar
);

insert into buildings (name, address)
values ('Дом 1', 'Липковского 37-Г'),
       ('Дом 2', 'Липковского 37-Б');