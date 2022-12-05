--ユーザーの作成
CREATE USER app_db;

--DBの作成
CREATE DATABASE app_db;

--ユーザーにDBの権限をまとめて付与
GRANT ALL PRIVILEGES ON DATABASE app_db TO app_db;

CREATE TABLE users (
    id BIGINT PRIMARY KEY,
    password VARCHAR(20) NOT NULL
);

INSERT INTO
    users
VALUES
    (1, 'password1');