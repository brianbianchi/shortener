package controllers

import (
	"encoding/json"
	"net/http"
	"shortener/models"
	"shortener/utils"
	"time"

	"github.com/gorilla/mux"
)

func Redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type,access-control-allow-origin, access-control-allow-headers")
	w.Header().Set("Access-Control-Expose-Headers", "Location")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	url := models.URL{}

	urlReceived, err := url.FindURLByCode(params["code"])
	if (*urlReceived == models.URL{}) {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	urlReceived.Visited++
	urlReceived.LastVisited = time.Now()
	err = url.IncrementURLVisitCount(urlReceived)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusSeeOther)
		return
	}
	http.Redirect(w, r, urlReceived.Link, http.StatusFound)
}

func GetUrls(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Content-Type", "application/json")
	url := models.URL{}

	urls, err := url.FindAllURLs()
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(urls)
}

func GetURLByCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	url := models.URL{}
	urlReceived, err := url.FindURLByCode(params["code"])
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	if (*urlReceived == models.URL{}) {
		http.Error(w, http.StatusText(404), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(urlReceived)
}

func CreateURL(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "http://localhost:8000")
	w.Header().Set("Content-Type", "application/json")
	var url models.URL
	_ = json.NewDecoder(r.Body).Decode(&url)

	// link validation
	if !utils.IsURL(url.Link) {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	// Unique link check
	urlFromLink, err := url.FindURLByLink(url.Link)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}
	if (*urlFromLink != models.URL{}) {
		json.NewEncoder(w).Encode(urlFromLink)
		return
	}

	// Unique code check
	for {
		url.Code = utils.RandSeq(6)
		urlFromCode, _ := url.FindURLByCode(url.Code)
		if (*urlFromCode == models.URL{}) {
			break
		}
	}

	url.Created = time.Now()
	url.Visited = 0
	url.LastVisited = time.Now()

	urlCreated, err := url.CreateURL(url)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(urlCreated)
}
