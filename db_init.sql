drop database xoho;
create database xoho;
use xoho;
drop table if exists areas;
create table areas(
    id int not null auto_increment primary key,
    name varchar(255),
    sub varchar(255));
drop table if exists prefectures;
create table prefectures(
    id int not null auto_increment primary key,
    name varchar(255),
    sub varchar(255));
drop table if exists types;
create table types(
    id int not null auto_increment primary key,
    name varchar(255),
    title varchar(255),
    sub varchar(255),
    opt varchar(255),
    icon_prefix varchar(255),
    icon_class varchar(255));
drop table if exists theaters;
create table theaters(
    id int not null auto_increment primary key,
    name varchar(255),
    sub varchar(255),
    area_id int,
    prefecture_id int,
    url varchar(255));
drop table if exists theater_types;
create table theater_types(
    theater_id int,
    type_id int);
