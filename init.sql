create database if not exists sitting_time_tracker;
use sitting_time_tracker;
create table if not exists users(
  uid int(10) primary key not null,
  password varchar(255) not null
);
create table if not exists records(
  uid int(10) not null,
  start_time datetime not null,
  end_time datetime not null,
  primary key (uid, start_time),
  foreign key (uid)
    references users(uid)
    on delete cascade
);