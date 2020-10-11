package models

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

func init() {
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
