package main

import (
	"fmt"
	"flag"
	"log"
	"os"

	"github.com/dlokkers/gophercises/link/parser"
)

func main() {
	htmlFilePtr := flag.String("file", "ex1.html", "html doc to parse")
	flag.Parse()

	htmlReader, err := os.Open(*htmlFilePtr)
	if err != nil {
		log.Fatal(err)
	}

	n, err := parse.FirstNode(htmlReader)
	if err != nil {
		log.Fatal(err)
	}
	links := parse.ParseHTML( n )
	fmt.Println(links)
}
