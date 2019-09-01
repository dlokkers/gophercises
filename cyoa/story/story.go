package cyoa

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"text/template"
)

type Story map[string]Chapter

type Chapter struct {
	Title      string `json:"title"`
	Paragraphs []string `json:"story"`
	Options    []Option
}

type Option struct {
	Text    string `json:"text"`
	Chapter string `json:"arc"`
}

type handler struct {
	s Story
}

func (h handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	tpl := template.Must(template.ParseFiles("template.html"))

	chapterName := strings.TrimPrefix(r.URL.Path, "/")

	if chapterName == "" {
		chapterName = "intro"
	} else if chapterName == "favicon.ico" {
		return
	}


	if chapter, ok := h.s[chapterName]; ok {
		err := tpl.Execute(w, chapter)
		if err != nil {
			log.Printf("%v", err)
			http.Error(w, "Something went wrong :(", http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Chapter not found", http.StatusNotFound)
}

func CreateStory(jsonBlob []byte) (Story, error) {
	var story Story
	err := json.Unmarshal(jsonBlob, &story)
	if err != nil {
		return nil, err
	}
	return story, nil
}

func NewHandler (s Story) http.Handler {
	return handler{s}
}
