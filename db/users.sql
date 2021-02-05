
create table users_db.users
(
    id         SERIAL PRIMARY KEY,
    first_name varchar(45),
    last_name  varchar(45),
    email      varchar(100) not null unique,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
);