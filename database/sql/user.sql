-- 生成用户表
create table if not exists account
( 
    id serial not null PRIMARY KEY,
    uid UUID  default gen_random_uuid(),
    name varchar(10) not null unique,
    email varchar(50) not null unique,
    phone varchar(11) unique,
    password varchar(255) not null,
    create_at timestamp default now(),
    update_at timestamp default now(),
    role_id int not null,
    avatar varchar(255) null
);
CREATE EXTENSION IF NOT EXISTS pgcrypto;

-- 插入管理员用户
insert into account (name, password, email,phone, role_id) values (
    'admin',
    crypt('1234', gen_salt('bf')),
    'admin@admin.com',
    '176025454xx',
1
);

-- 角色表
create table if not exists role
(
    id serial not null PRIMARY KEY,
    name varchar(30) not null unique,
    account_id int not null,
    casbin_role varchar(255) not null,
    create_at timestamp default now(),
    update_at timestamp default now()
);

-- 角色表初始化一条数据
insert into role (name, casbin_role, account_id) values (
    'admin',
    'admin',
    1
);