--ユーザーの作成
CREATE USER app_db;

--DBの作成
CREATE DATABASE app_db;

--ユーザーにDBの権限をまとめて付与
GRANT ALL PRIVILEGES ON DATABASE app_db TO app_db;

CREATE TABLE users (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    name NVARCHAR(255) NOT NULL,
    password VARCHAR(20) NOT NULL
);

CREATE TABLE invitation_users (
    id UUID NOT NULL DEFAULT gen_random_uuid(),
    invitation_code VARCHAR(100) NOT NULL,
    iv TEXT NOT NULL,
    key TEXT NOT NULL,
    encrypted_text TEXT NOT NULL,
    foreign key (id) references users(id)
);