drop table if exists "user" CASCADE ;
drop table if exists data;
create table if not exists "user"
(
    id        serial PRIMARY KEY,
    username  varchar(256),
    password  varchar(1024),
    telephone varchar(64),
    token     varchar(1024)
);
create table if not exists data
(
    id      serial PRIMARY KEY,
    title   varchar(256),
    number  int,
    content json,
    user_id int not null,
    foreign key (user_id) references "user" (id) on delete CASCADE
);
