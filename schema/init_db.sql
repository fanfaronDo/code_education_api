CREATE DATABASE notes;

CREATE TABLE notes(
    note_id INT PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    user_id INT NOT NULL
);

CREATE TABLE users(
    user_id INT PRIMARY KEY,
    name VARCHAR(255) NOT NULL ,
    username VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL
);