DROP DATABASE IF EXISTS rawati_test;

CREATE DATABASE IF NOT EXISTS rawati_test;

USE rawati_test;

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
);

CREATE TABLE user_profile (
    profile_id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    gender CHAR(1),
    birth_date DATE,
    height INT,
    weight INT,
    weight_goal INT,
    PRIMARY KEY (profile_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE,
    UNIQUE (user_id)
);

CREATE TABLE exercises (
    exercise_id INT AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    calories DECIMAL(6, 2) NOT NULL,
    PRIMARY KEY (exercise_id)
);

CREATE TABLE foods (
    food_id INT AUTO_INCREMENT,
    name VARCHAR(60) NOT NULL,
    calories DECIMAL(6, 2) NOT NULL,
    PRIMARY KEY (food_id)
);

CREATE TABLE exercise_per_day (
    exercise_activity_id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    name VARCHAR(60) NOT NULL,
    exercise_date DATE NOT NULL,
    duration INT NOT NULL,
    calories DECIMAL(6, 2),
    PRIMARY KEY (exercise_activity_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);

CREATE TABLE food_per_day (
    food_activity_id INT AUTO_INCREMENT,
    user_id INT NOT NULL,
    name VARCHAR(60) NOT NULL,
    food_date DATE NOT NULL,
    calories DECIMAL(6, 2) NOT NULL,
    PRIMARY KEY (food_activity_id),
    FOREIGN KEY (user_id) REFERENCES users(user_id) ON DELETE CASCADE
);
