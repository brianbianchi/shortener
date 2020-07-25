package main

import (
	"fmt"
	"log"
	"net/http"
)

func init() {
	dbConnect()
}

func main() {
	initializeRoutes()

	run(":8000")
}

func run(port string) {
	fmt.Printf("Listening to port %s\n", port)
	log.Fatal(http.ListenAndServe(port, router))
}
