package config

import (
	"os"
	"path/filepath"
	"encoding/json"
)

const configFileName = ".rsslaggconfig.json"

type Config struct {
	MaxPostsDisplayed	int			`json:"max_posts_displayed"`
	RSSFeedLinks		[]string	`json:"rss_feed_links"`
}
func Read() (Config, error) {
	fullPath, err := getConfigFilePath()	
	if err != nil {
		return Config{}, err
	}
	
	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}

func getConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	fullPath := filepath.Join(home, configFileName)
	return fullPath, nil
}
