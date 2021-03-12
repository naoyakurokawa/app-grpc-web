DROP SCHEMA IF EXISTS test;
CREATE SCHEMA test;
USE test;

DROP TABLE IF EXISTS users;
DROP TABLE IF EXISTS session;

CREATE TABLE users
(
  id int unsigned primary key auto_increment,
  name varchar(20),
  score int,
  photourl varchar(255),
  password varchar(255)
);

CREATE TABLE session
(
  id int unsigned primary key auto_increment,
  uuid varchar(255),
  name varchar(20),
  userid int
);

INSERT INTO users (name,score,photourl,password) VALUES ("test",100,"test.com","123");