package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
	rand.Seed(time.Now().UnixNano())
	db = InitDb()

	http.HandleFunc("/", RootHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("src/public"))))
	http.HandleFunc("/urls/", UrlHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Listening to port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
