CREATE SCHEMA test;

CREATE TABLE test.Person (
    ID serial PRIMARY KEY,
    Name VARCHAR(255) NOT NULL,
    LastUpdated TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP
);