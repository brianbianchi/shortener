package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

var (
	hostname     = os.Getenv("HOSTNAME")
	hostPort     = os.Getenv("HOSTPORT")
	username     = os.Getenv("USERNAME")
	password     = os.Getenv("DBPASSWORD")
	databaseName = os.Getenv("DBNAME")
)

var db *sql.DB

func dbConnect() {
	var err error
	// todo env vars
	connectionString := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable",
		username, password, hostname, hostPort, databaseName)

	db, err = sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}
}

func (u *URL) findAllURLs() (*[]URL, error) {
	rows, err := db.Query("SELECT * FROM urls")
	if err != nil {
		return &[]URL{}, err
	}
	defer rows.Close()

	urls := []URL{}

	for rows.Next() {
		url := URL{}
		err := rows.Scan(&url.Link, &url.Code, &url.Created, &url.Visited, &url.LastVisited)
		if err != nil {
			return &[]URL{}, err
		}
		urls = append(urls, url)
	}

	if err = rows.Err(); err != nil {
		return &[]URL{}, err
	}
	return &urls, nil
}

func (u *URL) findURLByCode(code string) (*URL, error) {
	rows, err := db.Query("SELECT * FROM urls WHERE code = $1", code)
	if err != nil {
		return &URL{}, err
	}
	defer rows.Close()

	url := URL{}

	for rows.Next() {
		err := rows.Scan(&url.Link, &url.Code, &url.Created, &url.Visited, &url.LastVisited)
		if err != nil {
			return &URL{}, err
		}
	}

	if err = rows.Err(); err != nil {
		return &URL{}, err
	}
	return &url, nil
}

func (u *URL) findURLByLink(link string) (*URL, error) {
	rows, err := db.Query("SELECT * FROM urls WHERE link = $1", link)
	if err != nil {
		return &URL{}, err
	}
	defer rows.Close()

	url := URL{}

	for rows.Next() {
		err := rows.Scan(&url.Link, &url.Code, &url.Created, &url.Visited, &url.LastVisited)
		if err != nil {
			return &URL{}, err
		}
	}

	if err = rows.Err(); err != nil {
		return &URL{}, err
	}
	return &url, nil
}

func (u *URL) createURL(url URL) (*URL, error) {
	rows, err := db.Query(`INSERT INTO urls (link, code, created, visited, last_visited) VALUES ($1, $2, $3, $4, $5)`,
		url.Link, url.Code, url.Created, url.Visited, url.LastVisited)
	if err != nil {
		return &URL{}, err
	}
	defer rows.Close()

	urlSaved, err := url.findURLByCode(url.Code)

	return urlSaved, err
}

func (u *URL) incrementURLVisitCount(url *URL) error {
	rows, err := db.Query(`Update urls SET visited = $1, last_visited = $2 WHERE code = $3`, url.Visited, url.LastVisited, url.Code)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
