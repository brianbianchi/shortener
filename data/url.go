package data

import (
	"database/sql"
	"fmt"
)

type URL struct {
	Link        string `json:"link"`
	Code        string `json:"code"`
	Created     string `json:"created"`
	Visited     int    `json:"visited"`
	LastVisited string `json:"last_visited"`
}

func GetURLs(db *sql.DB) (*[]URL, error) {
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

func GetURLByCode(db *sql.DB, code string) (*URL, error) {
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

func GetURLByLink(db *sql.DB, link string) (*URL, error) {
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

func CreateURL(db *sql.DB, url URL) error {
	_, err := db.Exec(`INSERT INTO urls (link, code, created, visited, last_visited) 
		VALUES ($1, $2, $3, $4, $5)`,
		url.Link, url.Code, url.Created, url.Visited, url.LastVisited)
	if err != nil {
		return err
	}

	return nil
}

func UpdateURL(db *sql.DB, url *URL) error {
	_, err := db.Exec(`Update urls 
		SET visited = $1, last_visited = $2 WHERE code = $3`,
		url.Visited, url.LastVisited, url.Code)
	if err != nil {
		return err
	}

	return nil
}
