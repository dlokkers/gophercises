package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/dlokkers/gophercises/quiet_hn/handlers"
	"github.com/dlokkers/gophercises/quiet_hn/hnapi"
)

func main() {
	var port, numStories, interval int
	flag.IntVar(&port, "port", 8080, "The port to start the web server on")
	flag.IntVar(&numStories, "num_stories", 30, "The number of top stories to display")
	flag.IntVar(&interval, "interval", 15, "The number of minutes between cache refreshes")
	flag.Parse()

	var cache hnapi.Cache
	var c hnapi.Client

	go cache.Refresh(c, numStories, time.Duration(interval)*time.Minute)

	http.HandleFunc("/", handlers.ShowStories(&cache))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
