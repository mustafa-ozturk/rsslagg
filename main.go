package main

import (
	"log"

	"github.com/mustafa-ozturk/rsslagg/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	rssItems, err := GetRSSItems(cfg.RSSFeedLinks)
	if err != nil {
		log.Fatalf("error getting RSS Items: %v", err)
	}
	rssItems = SortRSSItemsByDate(rssItems)
	PrintRSSItems(rssItems, cfg.MaxPostsDisplayed)
}
