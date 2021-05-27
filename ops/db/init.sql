CREATE DATABASE IF NOT EXISTS go_clean_arch;
USE go_clean_arch;

CREATE TABLE IF NOT EXISTS user (
    id varchar(50) PRIMARY KEY,
    email varchar(255) NOT NULL,
    password varchar(100) NOT NULL,
    name varchar(255) NOT NULL,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
);

CREATE TABLE IF NOT EXISTS book (
    id varchar(50) PRIMARY KEY,
    title varchar(255) NOT NULL,
    author varchar(255) NOT NULL,
    isbn varchar(255) NOT NULL,
    pages int NOT NULL,
    quantity int NOT NULL,
    created_at datetime,
    updated_at datetime,
    deleted_at datetime
);

CREATE TABLE IF NOT EXISTS book_user (
    user_id varchar(50),
    book_id varchar(50),
    PRIMARY KEY (`user_id`, `book_id`)
);