> API that generates a shorter, more shareable link.

## API

| Request method | Endpoint  | Description                                               |
| :------------- | :-------- | :-------------------------------------------------------- |
| GET            | `/`       | Returns static web interface from the `./public/` folder. |
| GET            | `/{code}` | Redirects to the associated link.                         |
| GET            | `/urls/`  | Returns all long URLs and related short codes.            |
| POST           | `/urls/`  | Creates new short URL from a `link`.                      |

## `URL` model

| Property       | Type      | Example                        | Description                                                                  |
| :------------- | :-------- | :----------------------------- | :--------------------------------------------------------------------------- |
| `link`         | `String`  | `"https://www.w3schools.com/"` | Long URL the user wants to redirect to.                                      |
| `code`         | `String`  | `"abcdeF"`                     | Server generated, 6-digit, short code used in the short URL. `[a-z,A-Z,0-9]` |
| `created`      | `Date`    | `2004-10-19T10:23:54Z`         | Date the short URL was created.                                              |
| `visited`      | `Integer` | `24`                           | Number of times the short URL was visited.                                   |
| `last_visited` | `Date`    | `"2010-10-19T10:23:54Z"`       | Date the short URL was last visited.                                         |

## Local setup

```console
$ brew install postgresql
$ psql postgres
$ postgres=>  CREATE ROLE app_user WITH LOGIN PASSWORD 'pw';
$ postgres=>  ALTER ROLE app_user CREATEDB;
$ postgres=>  \du  #list users
$ psql postgres -U app_user #enter as app_user
$ postgres=> CREATE DATABASE shortener;
$ postgres=>  \l #list dbs
$ ...
$ postgres=> \connect shortener
$ shortener=> \dl #list tables
```

> You can use the queries in `./sql/db.sql` to build the database with sample data.

## TODO

- tests
- copy short url after posting link
