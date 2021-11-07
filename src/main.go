package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

var (
	db *sql.DB
)

func main() {
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
