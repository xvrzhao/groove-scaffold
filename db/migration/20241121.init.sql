create table users (
  id serial primary key not null,
  username varchar(50) not null,
  nickname varchar(20) not null,
  phone varchar(20) null default null,
  email varchar(50) null default null,
  pwd_hash char(32) not null,
  salt char(8) not null,
  created_at timestamptz not null default now(),
  updated_at timestamptz not null default now(),
  deleted_at timestamptz null default null
);
create unique index users_username on users(username);