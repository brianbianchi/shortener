CREATE DATABASE shorten;
CREATE TABLE urls (
    link varchar(255) NOT NULL,
    code varchar(255) PRIMARY KEY NOT NULL,
    created TIMESTAMP NOT NULL,
    visited INTEGER NOT NULL,
    last_visited TIMESTAMP NOT NULL
);

INSERT INTO urls (link, code, created, visited, last_visited) 
VALUES ('https://www.w3schools.com/sql/sql_insert.asp', 'abcdeF', '2004-10-19 10:23:54', 1, '2010-10-19 10:23:54');
INSERT INTO urls (link, code, created, visited, last_visited) 
VALUES ('https://xkcd.com/2333/', 'fedcbA', '2004-10-19 10:23:54', 2, '2019-10-19 10:23:54');