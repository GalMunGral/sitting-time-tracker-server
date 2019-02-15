create database if not exists sitting_time_tracker;
use sitting_time_tracker;
create table if not exists users(
  uid int(10) primary key not null auto_increment,
  username varchar(255) not null unique,
  password varchar(255) not null
);
alter table users auto_increment=1000000000;
create table if not exists records(
  uid int(10) not null,
  start_time datetime not null,
  end_time datetime not null,
  primary key (uid, start_time),
  foreign key (uid)
    references users(uid)
    on delete cascade
);

# mysql://b9612c211f012c:6b9351b7@us-cdbr-iron-east-03.cleardb.net/heroku_dcd5cc19df4dc0e?reconnect=true