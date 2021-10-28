> API that generates a shorter, more shareable link.

## API

| Endpoint      | Description                                               |
| :------------ | :-------------------------------------------------------- |
| GET `/`       | Returns static web interface form the `./public/` folder. |
| GET `/{code}` | Redirects to the associated long URL.                     |
| GET `/urls`   | Returns all long URLs and related short codes.            |
| POST `/urls`  | Creates new short URL from a `link`.                      |

## `URL` model

```json
{
  "link": "https://www.w3schools.com/sql/sql_insert.asp",
  "code": "abcdeF",
  "created": "2004-10-19T10:23:54Z",
  "visited": 1,
  "last_visited": "2010-10-19T10:23:54Z"
}
```

## Local setup

```console
$ brew install postgresql
$ psql postgres
$ postgres=#  CREATE ROLE app_user WITH LOGIN PASSWORD 'pw';
$ postgres=#  ALTER ROLE app_user CREATEDB;
$ postgres=#  \du // list users
$ psql postgres -U app_user
$ postgres=> CREATE DATABASE shortener;
$ postgres=>  \l // list dbs


$ postgres=> \connect shortener
$ shortener=> \dl // list tables
```

## TODO

- tests
- copy short url after posting link
- display post error
- refresh url list when new posted
- list urls by most visited, 10 at a time, fixed table cell width
- html font
- URL model docs
