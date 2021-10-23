# shortener

API that generates a shorter, more shareable link.

## `GET` /api/redirect/{code}

Redirects to the associated long URL.

## `GET` /api/urls

Returns all long URLs and related short codes.

Response

```json
[
  {
    "link": "https://www.w3schools.com/sql/sql_insert.asp",
    "code": "abcdeF",
    "created": "2004-10-19T10:23:54Z",
    "visited": 1,
    "last_visited": "2010-10-19T10:23:54Z"
  },
  {
    "link": "http://youtube.com/",
    "code": "BpLnfg",
    "created": "2020-07-21T22:50:57.064527Z",
    "visited": 0,
    "last_visited": "2020-07-21T22:50:57.064528Z"
  }
]
```

## `GET` /api/urls/{code}

Returns all information related to a short code.

Response:

```json
{
  "link": "https://www.w3schools.com/sql/sql_insert.asp",
  "code": "abcdeF",
  "created": "2004-10-19T10:23:54Z",
  "visited": 1,
  "last_visited": "2010-10-19T10:23:54Z"
}
```

## `POST` /api/urls

Creates new short URL.

Body

```json
{
  "link": "http://youtube.com/ppoyijortujyio/khfdjhalkjfhlgskl?playback=20s"
}
```

Response

```json
{
  "link": "http://youtube.com/ppoyijortujyio/khfdjhalkjfhlgskl?playback=20s",
  "code": "jjJkwz",
  "created": "2020-12-17T17:16:20.506812Z",
  "visited": 0,
  "last_visited": "2020-12-17T17:16:20.506812Z"
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
