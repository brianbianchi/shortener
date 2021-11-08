package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

type APIResponseError struct {
	Status  int
	Message string
}

type APIResponseUrl struct {
	Status int
	Url    URL
}

type APIResponseUrls struct {
	Status int
	Urls   []URL
}

func RootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	key := strings.TrimPrefix(r.URL.Path, "/")

	if key != "" {
		urlReceived, err := FindURLByCode(key)
		if urlReceived == nil {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(APIResponseError{
				Status:  http.StatusNotFound,
				Message: fmt.Sprintf("Url with code %s not found.", key),
			})
			return
		}
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(APIResponseError{
				Status:  http.StatusInternalServerError,
				Message: "Server error.",
			})
			return
		}

		urlReceived.Visited++
		urlReceived.LastVisited = time.Now()

		err = IncrementURLVisitCount(urlReceived)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(APIResponseError{
				Status:  http.StatusInternalServerError,
				Message: "Server error. Failed to increment visited count.",
			})
			return
		}
		http.Redirect(w, r, urlReceived.Link, http.StatusFound)
		return
	}

	p := path.Dir("./src/public/index.html")
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
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(APIResponseError{
				Status:  http.StatusInternalServerError,
				Message: "Server error. Unable to find URLs.",
			})
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(APIResponseUrls{Status: http.StatusOK, Urls: *urls})
	case "POST":
		var url URL
		_ = json.NewDecoder(r.Body).Decode(&url)

		if !IsURL(url.Link) {
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(APIResponseError{Status: http.StatusBadRequest, Message: "Url is not a link."})
			return
		}

		existingUrl, _ := FindURLByLink(url.Link)
		if (*existingUrl != URL{}) {
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(APIResponseUrl{Status: http.StatusOK, Url: *existingUrl})
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
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(APIResponseError{
				Status:  http.StatusInternalServerError,
				Message: "Server error. Failed to create URL.",
			})
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(APIResponseUrl{
			Status: http.StatusCreated,
			Url:    *urlCreated,
		})
	}
}
