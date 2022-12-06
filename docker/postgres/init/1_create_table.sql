--ユーザーの作成
CREATE USER app_db;

--DBの作成
CREATE DATABASE app_db;

--ユーザーにDBの権限をまとめて付与
GRANT ALL PRIVILEGES ON DATABASE app_db TO app_db;

CREATE TABLE users (
    id UUID PRIMARY KEY NOT NULL COMMENT 'ユーザID',
    name VARCHAR(255) NOT NULL COMMENT 'ユーザ名',
);

CREATE TABLE invitation (
    id UUID PRIMARY KEY NOT NULL references users(id) COMMENT '招待コードを生成したユーザ',
    invitation_code VARCHAR(100) NOT NULL UNIQUE references invitation_codes(code),
    iv TEXT NOT NULL,
    key TEXT NOT NULL,
    encrypted_text TEXT NOT NULL
);

CREATE TABLE invitation_codes (
    code VARCHAR(100) PRIMARY KEY NOT NULL UNIQUE
);