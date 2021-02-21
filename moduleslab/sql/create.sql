DROP DATABASE go_rest_api;
CREATE DATABASE go_rest_api;
USE go_rest_api;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    email VARCHAR(50) NOT NULL,
    password VARCHAR(120) NOT NULL,
    age INT NOT NULL,
    active BOOL DEFAULT TRUE
);
CREATE UNIQUE INDEX uidx_email ON users (email);