package main

import (
	"fmt"
	"log"
	"context"
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

type RSSItemWithChannelTitle struct {
	ChannelTitle	string	
	ItemTitle		string	
	Link			string
	PubDate			time.Time
}

func main() {
	// 1. get feeds from config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Println("config:max_posts_displayed:", cfg.MaxPostsDisplayed)
	fmt.Println("config:rss_feed_links[0]:", cfg.RSSFeedLinks[0])
	fmt.Println("config:rss_feed_links[1]:", cfg.RSSFeedLinks[1])
	fmt.Println("config:rss_feed_links[2]:", cfg.RSSFeedLinks[2])

	rssItems := []RSSItemWithChannelTitle{}
	
	for _, link := range cfg.RSSFeedLinks {
		feed, err := fetchFeed(context.Background(), link)
		if err != nil {
			log.Fatalf("couldn't fetch feed: %v", err)
			return
		}

		for _, item := range feed.Channel.Item {
			pubDate, err := extractTimeFromPubDate(item.PubDate)
			if err != nil {
				log.Fatalf("$$ couldn't parse date: %v", err)
				return
			}

			rssItems = append(rssItems, RSSItemWithChannelTitle{
				ChannelTitle: feed.Channel.Title,
				ItemTitle: item.Title,
				Link:	item.Link,
				PubDate: pubDate,
			})
		}
	}


	for _, item := range rssItems {
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


	// 4. display latest max posts, default to 10 and ovewrite if -m flag present
}
