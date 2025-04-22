CREATE TABLE IF NOT EXISTS users (
    id serial PRIMARY KEY,
    name varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
);
