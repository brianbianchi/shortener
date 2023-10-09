package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strings"
	"text/template"
	"time"

	"github.com/brianbianchi/shortener/data"
)

type PageData struct {
	Urls   []data.URL
	NewUrl data.URL
	Error  string
}

func main() {
	db := data.InitDb()
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			key := strings.TrimPrefix(r.URL.Path, "/")

			if key != "" {
				urlReceived, err := data.GetURLByCode(db, key)
				if err != nil {
					http.Error(w, err.Error(), http.StatusBadRequest)
				}

				urlReceived.Visited++
				urlReceived.LastVisited = time.Now().String()
				_ = data.UpdateURL(db, urlReceived)

				http.Redirect(w, r, urlReceived.Link, http.StatusFound)
			}

			serveHome(w, db, data.URL{}, "")
		}

		if r.Method == "POST" {
			var url data.URL
			err := r.ParseForm()
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			url.Link = r.FormValue("link")

			if !data.IsURL(url.Link) {
				serveHome(w, db, data.URL{}, "You didn't submit a valid link.")
				return
			}

			existingUrl, err := data.GetURLByLink(db, url.Link)
			if err != nil {
				serveHome(w, db, data.URL{}, "Unable to look up link.")
				return
			}
			if existingUrl.Code != "" {
				serveHome(w, db, *existingUrl, "")
				return
			}

			for {
				url.Code = data.RandSeq(6)
				urlFromCode, err := data.GetURLByCode(db, url.Code)
				if err == nil && (*urlFromCode == data.URL{}) {
					break
				}
			}

			now := time.Now().Local().String()
			url.Created = now
			url.Visited = 0
			url.LastVisited = now

			err = data.CreateURL(db, url)
			if err != nil {
				serveHome(w, db, data.URL{}, "Unable to create a short URL.")
				return
			}
			serveHome(w, db, url, "")
		}
	})

	port := "3000"
	fmt.Println("Serving on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func serveHome(w http.ResponseWriter, db *sql.DB, newUrl data.URL, errorDisplay string) {
	path := data.GetRootPath()
	template, err := template.ParseFiles(fmt.Sprint(path, "index.html"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	urls, err := data.GetURLs(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	indexData := &PageData{
		Urls:   *urls,
		NewUrl: newUrl,
		Error:  errorDisplay,
	}
	err = template.Execute(w, indexData)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
