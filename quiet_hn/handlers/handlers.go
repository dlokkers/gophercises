package handlers

import (
	"html/template"
	"net/http"
	"time"

	"github.com/dlokkers/gophercises/quiet_hn/hnapi"
)

type templateData struct {
	Items      []hnapi.Item
	Time       time.Duration
	LastUpdate string
}

// ShowStories will load the top 30 stories to display
func ShowStories(c *hnapi.Cache) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tpl := template.Must(template.ParseFiles("handlers/index.gohtml"))
		start := time.Now()
		items := c.Items
		elapsed := time.Now().Sub(start)

		templateData := templateData{
			Items:      items,
			Time:       elapsed,
			LastUpdate: c.LastUpdate.Format("Mon Jan _2 15:04:05 2006"),
		}

		_ = tpl.Execute(w, templateData)
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
