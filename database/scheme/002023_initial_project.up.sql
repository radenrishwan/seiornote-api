create table users
(
    id         varchar(255) primary key,
    email      varchar(255) unique not null,
    username   varchar(255) unique not null,
    password   varchar(255)        not null,
    role       varchar(15)        not null,
    created_at varchar(255)        not null,
    updated_at varchar(255)        not null
);

create table sessions
(
    id         varchar(255) primary key,
    user_id    varchar(255) unique not null,
    created_at varchar(255)        not null,
    expired_at varchar(255)        not null,
    foreign key (user_id) references users (id) on delete cascade
);

create table notes
(
    id         varchar(255) primary key,
    user_id    varchar(255) not null,
    title      varchar(255) not null,
    content    text,
    favorite   bool         not null,
    created_at varchar(255) not null,
    updated_at varchar(255) not null,
    foreign key (user_id) references users (id) on delete cascade
);