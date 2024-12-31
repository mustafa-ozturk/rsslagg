package main

import (
	"log"
	"time"
	"strings"

	"github.com/mustafa-ozturk/rsslagg/internal/config"
)

func extractTimeFromPubDate(originalPubDate string) (time.Time, error) {
	splitDate := strings.Split(originalPubDate, " ")
	joinedDate := strings.Join(splitDate[:len(splitDate) - 2], " ")
	pubDate, err := time.Parse("Mon, 02 Jan 2006", joinedDate)
	return pubDate, err
}


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
