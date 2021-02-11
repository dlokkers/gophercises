package handlers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/dlokkers/gophercises/quiet_hn/hnapi"
)

type templateData struct {
	Items []hnapi.Item
	Time  time.Duration
}

// ShowStories will load the top 30 stories to display
func ShowStories() http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("handlers/index.html"))

		var c hnapi.Client
		topItems, err := c.GetTopItems()
		if err != nil {
			http.Error(w, "Failed to load top stories", http.StatusInternalServerError)
			return
		}

		var items []hnapi.Item
		start := time.Now()

		i := make(chan hnapi.Item)
		for _, id := range topItems {
			go retrieveStory(c, id, i)
		}
		for len(items) < 30 {
			items = append(items, <-i)
		}

		elapsed := time.Now().Sub(start)

		templateData := templateData{
			Items: items,
			Time:  elapsed,
		}
		err = tpl.Execute(w, templateData)
	})
}

func retrieveStory(c hnapi.Client, id int, i chan hnapi.Item) {
	item, err := c.GetItem(id)
	if err != nil {
		return
	}

	if !item.IsStory() {
		return
	}

	i <- item
}
