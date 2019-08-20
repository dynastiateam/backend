create table if not exists "user"
(
    id           serial
        constraint user_pk
            unique,
    appartment   int,
    email        varchar,
    password     varchar,
    phone        varchar,
    first_name   varchar,
    last_name    varchar,
    role         varchar,
    residence_id int,
    building_id  int
);

create index user_email_index
    on "user" (email);

