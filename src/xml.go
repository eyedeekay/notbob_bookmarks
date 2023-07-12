package nbjson

import (
	"log"
	"os"
	"strings"

	"github.com/mmcdole/gofeed"
)

func LoadXMLFile(filename string) (*gofeed.Feed, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	fp := gofeed.NewParser()
	feed, err := fp.Parse(file)
	return feed, err
}

type Bookmarks struct {
	Title string `json:"title"`
	URL   string `json:"url"`
}

func LoopOverXMLFile(feed *gofeed.Feed) []Bookmarks {
	bookmarks := []Bookmarks{}
	for _, item := range feed.Items {
		if strings.HasSuffix(item.Title, ".i2p.") {
			hostname := strings.Split(item.Title, " ")[len(strings.Split(item.Title, " "))-1]
			log.Println("Title:", "http://"+hostname)
			log.Println("Description:", item.Description)
			bookmarks = append(bookmarks, Bookmarks{
				Title: item.Title + " - " + item.Description,
				URL:   "http://" + hostname,
			})
		}
	}
	log.Println("XML Phase read finished")
	return bookmarks
}
