package models

import (
	"time"
)

// URL model
type URL struct {
	Link        string    `json:"link"`
	Code        string    `json:"code"`
	Created     time.Time `json:"created"`
	Visited     int       `json:"visited"`
	LastVisited time.Time `json:"last_visited"`
}

func (u *URL) FindAllURLs() (*[]URL, error) {
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

func (u *URL) FindURLByCode(code string) (*URL, error) {
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

func (u *URL) FindURLByLink(link string) (*URL, error) {
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

func (u *URL) CreateURL(url URL) (*URL, error) {
	rows, err := db.Query(`INSERT INTO urls (link, code, created, visited, last_visited) VALUES ($1, $2, $3, $4, $5)`,
		url.Link, url.Code, url.Created, url.Visited, url.LastVisited)
	if err != nil {
		return &URL{}, err
	}
	defer rows.Close()

	urlSaved, err := url.FindURLByCode(url.Code)

	return urlSaved, err
}

func (u *URL) IncrementURLVisitCount(url *URL) error {
	rows, err := db.Query(`Update urls SET visited = $1, last_visited = $2 WHERE code = $3`, url.Visited, url.LastVisited, url.Code)
	if err != nil {
		return err
	}
	defer rows.Close()

	return nil
}
