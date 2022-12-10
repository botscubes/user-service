CREATE TABLE account (
    id serial PRIMARY KEY,
    login VARCHAR ( 50 ) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    name VARCHAR (50) NOT NULL
);