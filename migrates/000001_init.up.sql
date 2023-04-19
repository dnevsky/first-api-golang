CREATE TABLE users
(
    id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    username VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL
);

CREATE TABLE todo_lists
(
    id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255)
);

CREATE TABLE users_lists
(
    id serial PRIMARY KEY,
    user_id int REFERENCES users(id) ON DELETE CASCADE NOT NULL,
    list_id int REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE todo_items
(
    id serial PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    done BOOLEAN NOT NULL DEFAULT FALSE
);

CREATE TABLE lists_items
(
    id serial PRIMARY KEY,
    item_id int REFERENCES todo_items(id) ON DELETE CASCADE NOT NULL,
    list_id int REFERENCES todo_lists(id) ON DELETE CASCADE NOT NULL
);