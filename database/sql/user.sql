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
    role varchar(30) not null,
    avatar varchar(255) null
);
-- 插入管理员用户
insert into account (name, password, email,phone, role) values (
    'admin',
    crypt('1234', gen_salt('bf')),
    'admin@admin.com',
    '176025454xx',
    'admin'
);