CREATE DATABASE IF NOT EXISTS codig267_golang;
USE codig267_golang;
DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(154) not null ,
    criadoEm timestamp default current_timestamp()
)ENGINE=INNODB;
