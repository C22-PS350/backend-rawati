DROP DATABASE IF EXISTS rawati;

CREATE DATABASE IF NOT EXISTS rawati;

USE rawati;

CREATE TABLE users (
    user_id INT AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    username VARCHAR(30) NOT NULL,
    email VARCHAR(60) NOT NULL,
    password CHAR(60) NOT NULL,
    is_verified BOOLEAN,
    PRIMARY KEY (user_id),
    UNIQUE (username),
    UNIQUE (email)
);

CREATE TABLE user_token (
    user_id INT,
    token CHAR(40),
    created_at DATETIME NOT NULL,
    PRIMARY KEY (user_id, token),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE (token)
)
