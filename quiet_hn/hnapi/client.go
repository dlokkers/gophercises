package hnapi

import (
	"encoding/json"
	"fmt"
	"net/http"
)

const (
	apiURL = "https://hacker-news.firebaseio.com/v0"
)

// Client is an API client to hacker news
type Client struct {
	apiURL string
}

// init will initialize the client if it's not ready for use yet
func (c *Client) init() {
	if c.apiURL == "" {
		c.apiURL = apiURL
	}
}

// GetTopItems returns the top items of HackerNews
func (c *Client) GetTopItems() ([]int, error) {
	c.init()
	resp, err := http.Get(fmt.Sprintf("%s/topstories.json", c.apiURL))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var topItems []int
	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&topItems)
	if err != nil {
		return nil, err
	}

	return topItems, nil
}

// GetItem will retrieve an individual HN item, by item ID
func (c *Client) GetItem(id int) (Item, error) {
	c.init()
	var item Item

	resp, err := http.Get(fmt.Sprintf("%s/item/%d.json", c.apiURL, id))
	if err != nil {
		return item, err
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	err = dec.Decode(&item)
	if err != nil {
		return item, err
	}

	return item, nil
}
