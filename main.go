package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"shortener/controllers"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

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
