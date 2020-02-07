DROP DATABASE go_rest_api;
CREATE DATABASE go_rest_api;
USE go_rest_api;
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    age INT NOT NULL
);