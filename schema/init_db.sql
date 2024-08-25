CREATE DATABASE notes;

\c notes

CREATE TABLE users(
    user_id SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password TEXT NOT NULL
);

CREATE TABLE notes(
    note_id SERIAL NOT NULL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    user_id INT REFERENCES users(user_id) ON DELETE CASCADE NOT NULL
);

