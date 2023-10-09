package main

import (
	"database/sql"
	"time"

	"github.com/brianbianchi/shortener/data"
)

func main() {
	db := data.InitDb()
	defer db.Close()

	CreateTables(db)
	CreateData(db)
}

func CreateTables(db *sql.DB) {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS urls (
			link TEXT NOT NULL,
			code TEXT PRIMARY KEY NOT NULL,
			created TEXT NOT NULL,
			visited INTEGER NOT NULL,
			last_visited TEXT NOT NULL
		);
	`)
	if err != nil {
		panic(err)
	}
}

func CreateData(db *sql.DB) {
	now := time.Now().Local().String()
	url := data.URL{
		Link:        "https://www.w3schools.com/sql/sql_insert.asp",
		Code:        "abcdeF",
		Created:     now,
		Visited:     1,
		LastVisited: now,
	}
	err := data.CreateURL(db, url)
	if err != nil {
		panic(err)
	}

	url.Link = "https://xkcd.com/2333/"
	url.Code = "fedcbA"
	url.Visited = 2
	err = data.CreateURL(db, url)
	if err != nil {
		panic(err)
	}
}
