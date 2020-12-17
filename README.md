# shortener
API that generates a shorter, more shareable link.

### endpoints
* `GET /api/redirect/{code}`
  * redirects to the associated long URL.
* `GET /api/urls`
  * returns all long URLs and related short codes.
  * response:
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
* `GET /api/urls/{code}`
  * returns all information related to a short code.
  * response:
```json
{
        "link": "https://www.w3schools.com/sql/sql_insert.asp",
        "code": "abcdeF",
        "created": "2004-10-19T10:23:54Z",
        "visited": 1,
        "last_visited": "2010-10-19T10:23:54Z"
}
```
* `POST /api/urls/{code}`
  * creates new transaction
  * body:
```json
{
	    "link": "http://youtube.com/ppoyijortujyio/khfdjhalkjfhlgskl?playback=20s"
}
```
  * response:
```json
{
        "link": "http://youtube.com/ppoyijortujyio/khfdjhalkjfhlgskl?playback=20s",
        "code": "jjJkwz",
        "created": "2020-12-17T17:16:20.506812Z",
        "visited": 0,
        "last_visited": "2020-12-17T17:16:20.506812Z"
}
```
