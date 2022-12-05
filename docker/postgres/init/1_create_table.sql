--ユーザーの作成
CREATE USER app_db;

--DBの作成
CREATE DATABASE app_db;

--ユーザーにDBの権限をまとめて付与
GRANT ALL PRIVILEGES ON DATABASE app_db TO app_db;

CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid() COMMENT 'ユーザID',
    name VARCHAR(255) NOT NULL COMMENT 'ユーザ名',
    password VARCHAR(20) NOT NULL
);

CREATE TABLE invitation_users (
    id UUID PRIMARY KEY NOT NULL DEFAULT gen_random_uuid() references users(id),
    invitation_code VARCHAR(100) NOT NULL UNIQUE,
    iv TEXT NOT NULL,
    key TEXT NOT NULL,
    encrypted_text TEXT NOT NULL,
);