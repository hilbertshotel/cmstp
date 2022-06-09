CREATE DATABASE asd;

\c asd

CREATE TABLE users (
  id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name VARCHAR (100) NOT NULL 
);

INSERT INTO users (name) VALUES ('Charlie');
INSERT INTO users (name) VALUES ('Bobbie');
INSERT INTO users (name) VALUES ('Merry');
