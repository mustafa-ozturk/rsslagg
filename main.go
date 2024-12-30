package main

import (
	"fmt"
	"log"
	"time"
	"strings"
	"sort"

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

	rssItems, err := GetRssItems(cfg.RSSFeedLinks)
	
	// sortRssItemsByDate
	// printRssItems

	// sort items by date
	sort.Slice(rssItems, func(i, j int) bool {
		return rssItems[i].PubDate.Before(rssItems[j].PubDate)
	})

	// print items
	for _, item := range rssItems[len(rssItems) - cfg.MaxPostsDisplayed:] {
		pubDateStr := fmt.Sprintf("%04d-%02d-%02d",
			item.PubDate.Year(),
			int(item.PubDate.Month()),
			item.PubDate.Day())

		fmt.Printf("- %s | %s | %s:\n\t%s\n\n",
			pubDateStr,
			item.ChannelTitle,
			item.ItemTitle,
			item.Link) 
	}
}
