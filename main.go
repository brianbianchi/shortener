package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path"
	"shortener/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	// http.Handle("/", http.FileServer(http.Dir("./public"))
	s := http.StripPrefix("/public/", http.FileServer(http.Dir("./public/")))
	router.PathPrefix("/public/").Handler(s)
	router.HandleFunc("/", home).Methods("GET")
	router.HandleFunc("/api/redirect/{code}", controllers.Redirect).Methods("GET")
	router.HandleFunc("/api/urls", controllers.GetUrls).Methods("GET")
	router.HandleFunc("/api/urls/{code}", controllers.GetURLByCode).Methods("GET")
	router.HandleFunc("/api/urls", controllers.CreateURL).Methods("POST")

	port := os.Getenv("PORT")
	if port == "" {
		port = "8000" //localhost
	}

	fmt.Printf("Listening to port %s\n", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}

func home(w http.ResponseWriter, r *http.Request) {
	p := path.Dir("./public/index.html")
	// set header
	w.Header().Set("Content-type", "text/html")
	http.ServeFile(w, r, p)
}
