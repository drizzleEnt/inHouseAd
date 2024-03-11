-- +goose Up
CREATE TABLE users(
    id SERIAL PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    password VARCHAR(255) NOT NULL
);
CREATE TABLE categories(
    id SERIAL PRIMARY KEY,
    categoryName VARCHAR(255) NOT NULL
);

CREATE TABLE goods(
    id SERIAL PRIMARY KEY,
    catigoryId INT[],
    goodsName VARCHAR(255) NOT NULL
);

-- +goose Down
DROP TABLE goods;
DROP TABLE categories;
DROP TABLE users;
