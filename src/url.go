package main

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

type URL struct {
	Link        string    `json:"link"`
	Code        string    `json:"code"`
	Created     time.Time `json:"created"`
	Visited     int       `json:"visited"`
	LastVisited time.Time `json:"last_visited"`
}

func FindAllURLs() (*[]URL, error) {
	rows, err := db.Query("SELECT * FROM urls ORDER BY last_visited DESC LIMIT 20")
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	defer rows.Close()

	urls := []URL{}

	for rows.Next() {
		url := URL{}
		err := rows.Scan(&url.Link, &url.Code, &url.Created, &url.Visited, &url.LastVisited)
		if err != nil {
			return nil, err
		}
		urls = append(urls, url)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &urls, nil
}

func FindURLByCode(code string) (*URL, error) {
	rows, err := db.Query("SELECT * FROM urls WHERE code = $1", code)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	url := URL{}

	for rows.Next() {
		err := rows.Scan(&url.Link, &url.Code, &url.Created, &url.Visited, &url.LastVisited)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &url, nil
}

func FindURLByLink(link string) (*URL, error) {
	rows, err := db.Query("SELECT * FROM urls WHERE link = $1", link)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	url := URL{}

	for rows.Next() {
		err := rows.Scan(&url.Link, &url.Code, &url.Created, &url.Visited, &url.LastVisited)
		if err != nil {
			return nil, err
		}
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return &url, nil
}

func CreateURL(url URL) (*URL, error) {
	rows, err := db.Query(`INSERT INTO urls (link, code, created, visited, last_visited) VALUES ($1, $2, $3, $4, $5)`,
		url.Link, url.Code, url.Created, url.Visited, url.LastVisited)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	urlSaved, err := FindURLByCode(url.Code)

	return urlSaved, err
}

func IncrementURLVisitCount(url *URL) error {
	rows, err := db.Query(`Update urls SET visited = $1, last_visited = $2 WHERE code = $3`, url.Visited, url.LastVisited, url.Code)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
