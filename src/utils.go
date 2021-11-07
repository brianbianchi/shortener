package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"net/url"
	"os"

	_ "github.com/lib/pq"
)

var (
	username     = os.Getenv("USERNAME")
	databaseName = os.Getenv("DBNAME")
	letters      = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func InitDb() *sql.DB {
	connectionString := fmt.Sprintf("user=%s dbname=%s sslmode=disable", username, databaseName)

	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		panic(err)
	}

	if err = db.Ping(); err != nil {
		panic(err)
	}

	return db
}

// Creates a random sequence of characters for the URL code
func RandSeq(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

// Validates that a string has a URL structure
func IsURL(str string) bool {
	u, err := url.Parse(str)
	return err == nil && u.Scheme != "" && u.Host != ""
}
