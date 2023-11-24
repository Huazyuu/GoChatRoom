drop database if exists go_chat;
create database go_chat;
use go_chat;
drop table  if exists chat_logs;
create table chat_logs
(
    `id`      bigint        not null auto_increment,
    `message` varchar(1000) null,
    `addtime` datetime default current_timestamp,
    `address` varchar(50)   null,
    primary key (`id`)
);