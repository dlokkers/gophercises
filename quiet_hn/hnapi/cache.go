package hnapi

import (
	"sort"
	"time"
)

// Cache stores all stories for display purposes
type Cache struct {
	Items      []Item
	Error      error
	LastUpdate time.Time
}

// Refresh refreshes the cache ever interval using client c
func (ch *Cache) Refresh(c Client, size int, interval time.Duration) {
	for {
		ch.Error = nil

		itemIDs, err := c.GetTopItems()
		if err != nil {
			ch.Error = err

			// If we have an error, retry faster than interval
			time.Sleep(10 * time.Second)
			continue
		}

		type retrieval struct {
			Index int
			Item  Item
			Error error
		}
		itemChannel := make(chan retrieval)

		for i := 0; i < size; i++ {
			go func(i, id int) {
				item, err := c.GetItem(id)
				ret := retrieval{
					Index: i,
					Item:  item,
					Error: err,
				}
				itemChannel <- ret
			}(i, itemIDs[i])
		}

		var retrievedItems []retrieval
		for i := 0; i < size; i++ {
			ret := <-itemChannel
			if ret.Item.IsStory() && ret.Error == nil {
				retrievedItems = append(retrievedItems, ret)
			}
		}

		sort.Slice(retrievedItems, func(i, j int) bool {
			return retrievedItems[i].Index < retrievedItems[j].Index
		})

		ch.Items = nil
		for _, ret := range retrievedItems {
			ch.Items = append(ch.Items, ret.Item)
		}
		ch.LastUpdate = time.Now()

		time.Sleep(interval)
	}
}
