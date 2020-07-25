package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func redirect(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	params := mux.Vars(r)
	url := URL{}

	urlReceived, err := url.findURLByCode(params["code"])
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
	err = url.incrementURLVisitCount(urlReceived)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, urlReceived.Link, http.StatusSeeOther)
}

func getUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	url := URL{}

	urls, err := url.findAllURLs()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(urls)
}

func getURLByCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	url := URL{}
	urlReceived, err := url.findURLByCode(params["code"])
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	if (*urlReceived == URL{}) {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(urlReceived)
}

func createURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var url URL
	_ = json.NewDecoder(r.Body).Decode(&url)

	// link validation
	if !isURL(url.Link) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Unique link check
	urlFromLink, err := url.findURLByLink(url.Link)
	if (*urlFromLink != URL{}) {
		json.NewEncoder(w).Encode(urlFromLink)
		return
	}

	// Unique code check
	for {
		url.Code = randSeq(6)
		urlFromCode, _ := url.findURLByCode(url.Code)
		if (*urlFromCode == URL{}) {
			break
		}
	}

	url.Created = time.Now()
	url.Visited = 0
	url.LastVisited = time.Now()

	urlCreated, err := url.createURL(url)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(urlCreated)
}
