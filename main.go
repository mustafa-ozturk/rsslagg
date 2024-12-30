package main

import (
	"fmt"
	"log"
	"context"
	"time"
	"strings"

	"github.com/mustafa-ozturk/rsslagg/internal/config"
)

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

	rssFeeds := []RSSFeed{}
	
	for _, link := range cfg.RSSFeedLinks {
		feed, err := fetchFeed(context.Background(), link)
		if err != nil {
			log.Fatalf("couldn't fetch feed: %v", err)
			return
		}

		rssFeeds = append(rssFeeds, *feed)
	}


	for _, feed := range rssFeeds {
		for _, item := range feed.Channel.Item {
			splitDate := strings.Split(item.PubDate, " ")
			joinedDate := strings.Join(splitDate[:len(splitDate) - 2], " ")
			pubDate, err := time.Parse("Mon, 02 Jan 2006", joinedDate)
			if err != nil {
				log.Fatalf("couldn't parse date: %v", err)
				return
			}

			pubDateStr := fmt.Sprintf("%04d-%02d-%02d",
				pubDate.Year(),
				int(pubDate.Month()),
				pubDate.Day())

			fmt.Printf("- %s | %s | %s:\n\t%s\n\n",
				pubDateStr,
				feed.Channel.Title,
				item.Title,
				item.Link) 
		}
	}


	// 3. sort by latest
	// 4. display latest max posts, default to 10 and ovewrite if -m flag present
}
