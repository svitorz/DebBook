CREATE DATABASE IF NOT EXISTS devbook;

USE devbook;

DROP TABLE IF EXISTS USUARIOS;

CREATE TABLE USUARIOS(
  id int auto_increment primary key,
  nome varchar(50) not null, 
  nick varchar(50) not null unique,
  email varchar(50) not null unique,
  senha varchar(50) not null, 
  criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;

DROP TABLE IF EXISTS SEGUIDORES;

CREATE TABLE SEGUIDORES(
  usuario_id int not null,
  FOREIGN KEY (usuario_id)
  REFERENCES USUARIOS(id)
  ON DELETE CASCADE,
  seguidor_id int not null,
  FOREIGN KEY (usuario_id)
  REFERENCES USUARIOS(id)
  ON DELETE CASCADE,

  PRIMARY KEY(usuario_id, seguidor_id)
) ENGINE=INNODB;
