package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dlokkers/gophercises/cyoa/story"
)


func main() {
	portPtr     := flag.Int("port", 3000, "Port to run the web app")
	fileNamePtr := flag.String( "file", "gopher.json", "the JSON file with the CYOA story")
	flag.Parse()

	jsonBlob, err := ioutil.ReadFile(*fileNamePtr)
	if err != nil {
		log.Fatal(err)
	}

	story, err := cyoa.CreateStory(jsonBlob)
	if err != nil {
		log.Fatal(err)
	}

	h := cyoa.NewHandler(story)
	fmt.Printf("Starting the server on port: %d\n", *portPtr)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", *portPtr),h))

	// display the first arc
}
