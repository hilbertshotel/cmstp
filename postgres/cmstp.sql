CREATE DATABASE asd;

\c asd

CREATE TABLE users (
  id BIGINT PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  name VARCHAR (100) NOT NULL 
);

INSERT INTO users (name) VALUES ('boce');
INSERT INTO users (name) VALUES ('koce');
INSERT INTO users (name) VALUES ('unknown');
INSERT INTO users (name) VALUES ('kolu');
