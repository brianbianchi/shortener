package main

import "github.com/gorilla/mux"

var router *mux.Router

func initializeRoutes() {
	router = mux.NewRouter()

	router.HandleFunc("/api/redirect/{code}", redirect).Methods("GET")
	router.HandleFunc("/api/urls", getUrls).Methods("GET")
	router.HandleFunc("/api/urls/{code}", getURLByCode).Methods("GET")

	router.HandleFunc("/api/urls", createURL).Methods("POST")
}
