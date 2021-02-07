package main

import (
	"fmt"

	"github.com/dlokkers/gophercises/quiet_hn/hnapi"
)

func main() {
	var c hnapi.Client
	topItems, err := c.GetTopItems()
	if err != nil {
		panic(err)
	}
	count := 0
	for _, id := range topItems {
		item, err := c.GetItem(id)
		if err != nil {
			panic(err)
		}
		if item.IsStory() {
			count++
			fmt.Printf("%d. %s\n", count, item.Title)
			fmt.Printf("-- %s\n", item.URL)
		}
		if count >= 30 {
			break
		}
	}
}
