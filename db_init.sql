create database xoho;
use xoho;
drop table if exists areas;
create table areas(
    id int not null auto_increment primary key,
    name varchar(255),
    sub varchar(255));
INSERT INTO areas (name, sub) VALUES ("東北地区", "TOHOKU AREA");
INSERT INTO areas (name, sub) VALUES ("関東地区", "KANTO AREA");
INSERT INTO areas (name, sub) VALUES ("中部地区", "CHUBU AREA");
INSERT INTO areas (name, sub) VALUES ("関西地区", "KANSAI AREA");
INSERT INTO areas (name, sub) VALUES ("中国地方", "CHUGOKU AREA");
INSERT INTO areas (name, sub) VALUES ("四国地区", "SHIKOKU AREA");
INSERT INTO areas (name, sub) VALUES ("九州地区", "KYUSYU AREA");
