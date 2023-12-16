# URL shortening service

## Purpose
- "beautify" a link
- track clicks
- disguise the underlying address

## Cons

- disguised link could be malicious
- some shortening service providers are blacklisted or considered to be spam
- some websites prevent short, redirected URLs from being posted

## `URL` schema

| Property       | Type      | Example                        | Description                                                                  |
| :------------- | :-------- | :----------------------------- | :--------------------------------------------------------------------------- |
| `link`         | `String`  | `"https://www.w3schools.com/"` | Long URL the user wants to redirect to.                                      |
| `code`         | `String`  | `"abcdeF"`                     | Server generated, 6-digit, short code used in the short URL. `[a-z,A-Z,0-9]` |
| `created`      | `String`    | `"2004-10-19T10:23:54Z"`         | Date the short URL was created.                                              |
| `visited`      | `Integer` | `24`                           | Number of times the short URL was visited.                                   |
| `last_visited` | `String`    | `"2010-10-19T10:23:54Z"`       | Date the short URL was last visited.                                         |
