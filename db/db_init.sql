mysql -u root -p --local-infile=1
set global local_infile=on;
create database xoho;
use xoho;
drop table if exists areas;
create table areas(
    id int not null auto_increment primary key,
    name varchar(255),
    sub varchar(255));
load data local
infile "areas.csv"
into table areas
fields terminated by ','
optionally enclosed by '"';
drop table if exists prefectures;
create table prefectures(
    id int not null auto_increment primary key,
    name varchar(255),
    sub varchar(255));
load data local
infile "prefectures.csv"
into table prefectures
fields terminated by ','
optionally enclosed by '"';
drop table if exists types;
create table types(
    id int not null auto_increment primary key,
    name varchar(255),
    title varchar(255),
    sub varchar(255),
    opt varchar(255),
    icon_prefix varchar(255),
    icon_class varchar(255));
load data local
infile "types.csv"
into table types
fields terminated by ','
optionally enclosed by '"'
ignore 1 lines
(
@name,
@title,
@sub,
@opt,
@icon_prefix,
@icon_class
)
set
name=@name,
title=@title,
sub=@sub,
opt=@opt,
icon_prefix=@icon_prefix,
icon_class=@icon_class
;
drop table if exists theaters;
create table theaters(
    id int not null auto_increment primary key,
    name varchar(255),
    sub varchar(255),
    area_id int,
    prefecture_id int,
    url varchar(255));
load data local
infile "theaters.csv"
into table theaters
fields terminated by ','
optionally enclosed by '"'
ignore 1 lines
(
@name,
@sub,
@area,
@prefecture,
@url
)
set
name=@name,
sub=@sub,
area_id=@area,
prefecture_id=@prefecture,
url=@url
;
drop table if exists theater_types;
create table theater_types(
    theater_id int,
    type_id int);
load data local
infile "theater_types.csv"
into table theater_types
fields terminated by ','
optionally enclosed by '"'
ignore 1 lines
(
@type_id,
@theater_id
)
set
type_id=@type_id,
theater_id=@theater_id
;
drop table if exists users;
create table users(
    id int not null auto_increment primary key,
    name varchar(255),
    password varchar(255),
    user_ext_id int
);
drop table if exists user_exts;
create table user_exts(
    id int not null auto_increment primary key,
    auth_miss_count int
);
