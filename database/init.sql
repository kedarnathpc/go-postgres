CREATE DATABASE stocksdb;

-- Grant privileges to the user on the database
GRANT ALL PRIVILEGES ON DATABASE stocksdb TO postgres;

-- Switch to the new database
\c stocksdb;

-- Create a table
CREATE TABLE stocks (
    stockid SERIAL NOT NULL PRIMARY KEY,
    name TEXT,
    price INT,
    company TEXT
);  

