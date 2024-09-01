CREATE DATABASE todo;

\c todo

CREATE TABLE users(
    id SERIAL UNIQUE,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL,
    password TEXT NOT NULL
);

CREATE TABLE todo_items(
    id    SERIAL UNIQUE,
    title VARCHAR(255) NOT NULL,
    description TEXT,
    due_date TIMESTAMP NOT NULL,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE user_todo_items(
    id SERIAL UNIQUE,
    user_id INT REFERENCES users(id) NOT NULL,
    todo_id INT REFERENCES todo_items(id) NOT NULL
);


