# syntax=docker/dockerfile:1

FROM mysql

ENV MYSQL_ROOT_PASSWORD astropass

EXPOSE 3306

# CMD mysql -v
# CMD mysql -uroot --password=astropass -e "create database AstroRanking; use AstroRanking; create table userRanking (id varchar(32), username varchar(32), timeInSeconds int, map int)"