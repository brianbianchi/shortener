package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"
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

type APIError struct {
	Status  int
	Message string
}

var (
	db           *sql.DB
	username     = os.Getenv("USERNAME")
	databaseName = os.Getenv("DBNAME")
	letters      = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
)

func DbInit() *sql.DB {
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

func FindAllURLs() (*[]URL, error) {
	rows, err := db.Query("SELECT * FROM urls")
	if err != nil {
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
		return &[]URL{}, err
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

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Content-Type", "application/json")
	key := strings.TrimPrefix(r.URL.Path, "/")

	// Redirect
	if key != "" {
		urlReceived, err := FindURLByCode(key)
		if (*urlReceived == URL{}) {
			http.Error(w, http.StatusText(404), http.StatusNotFound)
			return
		}
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		urlReceived.Visited++
		urlReceived.LastVisited = time.Now()
		err = IncrementURLVisitCount(urlReceived)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusSeeOther)
			return
		}
		http.Redirect(w, r, urlReceived.Link, http.StatusFound)
		return
	}

	p := path.Dir("./public/index.html")
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}

func UrlHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		urls, err := FindAllURLs()
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(201)
		json.NewEncoder(w).Encode(urls)
	case "POST":
		var url URL
		_ = json.NewDecoder(r.Body).Decode(&url)

		if !IsURL(url.Link) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIError{Status: http.StatusBadRequest, Message: "Url is not a link."})
			return
		}

		existingUrl, _ := FindURLByLink(url.Link)
		if (*existingUrl != URL{}) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(existingUrl)
			return
		}

		for {
			url.Code = RandSeq(6)
			urlFromCode, _ := FindURLByCode(url.Code)
			if (*urlFromCode == URL{}) {
				break
			}
		}

		url.Created = time.Now()
		url.Visited = 0
		url.LastVisited = time.Now()

		urlCreated, err := CreateURL(url)
		if err != nil {
			http.Error(w, http.StatusText(500), http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(urlCreated)
	}
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

func main() {
	db = DbInit()

	http.HandleFunc("/", RootHandler)
	http.Handle("/public/", http.StripPrefix("/public/", http.FileServer(http.Dir("public"))))
	http.HandleFunc("/urls/", UrlHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}

	fmt.Printf("Listening to port %s\n", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
