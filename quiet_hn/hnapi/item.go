package hnapi

// Item is a Hackernews item returned by the API
type Item struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	By          string `json:"by"`
	Text        string `json:"text"`
	Kids        []int  `json:"kids"`
	URL         string `json:"url"`
	Title       string `json:"title"`
	Descendants int    `json:"descendants"`
}

// IsStory will return true if an Item is a story
func (i *Item) IsStory() bool {
	if i.Type != "story" {
		return false
	}

	// Articles with text instead of URL are likely discussions
	if i.URL == "" {
		return false
	}

	return true
}
