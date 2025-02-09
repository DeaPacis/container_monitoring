package utils

import (
	"os"
	"strconv"
	"time"
)

const DefaultScrapeInterval = 5

func GetScrapeInterval() time.Duration {
	scrapeInterval, _ := strconv.Atoi(os.Getenv("SCRAPE_INTERVAL"))
	if scrapeInterval == 0 {
		scrapeInterval = DefaultScrapeInterval
	}
	return time.Duration(scrapeInterval) * time.Second
}
