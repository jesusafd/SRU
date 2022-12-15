CREATE DATABASE users;

USE users;

CREATE TABLE user (
    email varchar(255) not null primary key,
    password varchar(255) not null,
    name varchar(255) not null,
    last_name varchar(255) not null,
    year_birth int,
    month_birth int,
    day_birth int
);

