package main

import (
	"log"
	"net/http"

	"github.com/dlokkers/gophercises/quiet_hn/handlers"
)

func main() {
	http.HandleFunc("/", handlers.ShowStories())
	log.Fatal(http.ListenAndServe(":8080", nil))
}
