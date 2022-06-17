CREATE DATABASE test;

\c test

CREATE TABLE test (
  data TEXT
);

INSERT INTO test (data) VALUES ('some test string');
