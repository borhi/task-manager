CREATE TABLE IF NOT EXISTS project(
   id serial PRIMARY KEY,
   name VARCHAR (500) UNIQUE NOT NULL,
   description TEXT
);