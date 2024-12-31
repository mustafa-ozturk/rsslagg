package main

import (
	"context"
	"time"
	"net/http"
	"io"
	"encoding/xml"
	"html"
	"sort"
	"fmt"
	"strings"
)

type RSSFeed struct {
	Channel struct {
		Title		string		`xml:"title"`
		Link		string		`xml:"link"`
		Description	string		`xml:"description"`
		Item		[]RSSItem	`xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title		string	`xml:"title"`
	Link		string	`xml:"link"`
	Description	string	`xml:"description"`
	PubDate		string	`xml:"pubDate"`
}

type RSSItemWithChannelTitle struct {
	ChannelTitle	string	
	ItemTitle		string	
	Link			string
	PubDate			time.Time
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "rsslagg")
	resp, err := httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var rssFeed RSSFeed
	err = xml.Unmarshal(data, &rssFeed)
	if err != nil {
		return nil, err
	}

	rssFeed.Channel.Title = html.UnescapeString(rssFeed.Channel.Title)
	rssFeed.Channel.Description = html.UnescapeString(rssFeed.Channel.Description)
	for i, item := range rssFeed.Channel.Item {
		item.Title = html.UnescapeString(item.Title)
		item.Description = html.UnescapeString(item.Description)
		rssFeed.Channel.Item[i] = item
	}

	return &rssFeed, nil
}


func GetRSSItems(feedLinks []string) ([]RSSItemWithChannelTitle, error) {
	rssItems := []RSSItemWithChannelTitle{}
	for _, link := range feedLinks {
		feed, err := fetchFeed(context.Background(), link)
		if err != nil {
			return []RSSItemWithChannelTitle{}, err
		}

		for _, item := range feed.Channel.Item {
			pubDate, err := extractTimeFromPubDate(item.PubDate)
			if err != nil {
				return []RSSItemWithChannelTitle{}, err
			}

			rssItems = append(rssItems, RSSItemWithChannelTitle{
				ChannelTitle: feed.Channel.Title,
				ItemTitle: item.Title,
				Link:	item.Link,
				PubDate: pubDate,
			})
		}
	}

	return rssItems, nil
}

func SortRSSItemsByDate(rssItems []RSSItemWithChannelTitle) []RSSItemWithChannelTitle {
	sort.Slice(rssItems, func(i, j int) bool {
		return rssItems[i].PubDate.Before(rssItems[j].PubDate)
	})
	return rssItems
}

func PrintRSSItems(rssItems []RSSItemWithChannelTitle, maxPostsDisplayed int) {
	for _, item := range rssItems[len(rssItems) - maxPostsDisplayed:] {
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

func extractTimeFromPubDate(originalPubDate string) (time.Time, error) {
	splitDate := strings.Split(originalPubDate, " ")
	joinedDate := strings.Join(splitDate[:len(splitDate) - 2], " ")
	pubDate, err := time.Parse("Mon, 02 Jan 2006", joinedDate)
	return pubDate, err
}
