
create table users_db.users
(
    id         SERIAL PRIMARY KEY,
    first_name varchar(45),
    last_name  varchar(45),
    email      varchar(100) not null unique,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);

ALTER TABLE users_db.users
    ADD column status varchar(20) not null default '',
    ADD column pasword varchar(30) not null default '';