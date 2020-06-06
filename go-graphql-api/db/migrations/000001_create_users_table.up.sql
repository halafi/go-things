CREATE TABLE users (
  id serial PRIMARY KEY,
  name VARCHAR (50) NOT NULL,
  age INT NOT NULL,
  profession VARCHAR (50) NOT NULL,
  friendly BOOLEAN NOT NULL
);
