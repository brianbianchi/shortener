package data

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func InitDb() *sql.DB {
	path := GetRootPath()
	db, err := sql.Open("sqlite3", fmt.Sprint(path, "shortener.db"))
	if err != nil {
		panic(err)
	}
	return db
}
