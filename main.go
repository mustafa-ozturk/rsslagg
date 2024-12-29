package main

import (
	"fmt"
	"log"

	"github.com/mustafa-ozturk/rsslagg/internal/config"
)

func main() {
	// 1. get feeds from config file
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config: %v", err)
	}

	fmt.Println("config:max_posts_displayed:", cfg.MaxPostsDisplayed)
	fmt.Println("config:rss_feeds[0]:", cfg.RSSFeeds[0])
	fmt.Println("config:rss_feeds[1]:", cfg.RSSFeeds[1])
	fmt.Println("config:rss_feeds[2]:", cfg.RSSFeeds[2])


	// 2. get max post from config file
	// 3. sort by latest
	// 4. display latest max posts, default to 10 and ovewrite if -m flag present
}
