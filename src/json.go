package nbjson

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"time"
)

type JSONFeed struct {
	Directory string      `json:"directory"`
	Bookmarks []Bookmarks `json:"bookmarks"`
}

func BookmarksToJSONFeed(dirname string, bookmarks []Bookmarks) JSONFeed {
	for _, bookmark := range bookmarks {
		log.Println("Title:", bookmark.Title, "URL:", bookmark.URL)
	}
	feed := JSONFeed{
		Directory: dirname,
		Bookmarks: bookmarks,
	}
	return feed
}

func PrettyPrintJSONFeed(feed JSONFeed) []byte {
	for _, bookmark := range feed.Bookmarks {
		log.Println("Title:", bookmark.Title, "URL:", bookmark.URL)
	}
	bytes, err := json.MarshalIndent(feed, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(bytes))
	return bytes
}

func SaveJSONFeed(bytes []byte) {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	filename := fmt.Sprintf("notbob%d.json", timestamp)
	ioutil.WriteFile(filename, bytes, 0644)
}

func CombineAllFeeds() []byte {
	// Get a list of all .json files in the directory
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	// Create a set to store unique feed items
	//uniqueItems := make(map[string]bool)

	// Iterate over each file
	for _, file := range files {
		// Check if the file is a .json file
		if strings.HasSuffix(file.Name(), ".json") {
			// Read the contents of the file
			content, err := ioutil.ReadFile(file.Name())
			if err != nil {
				log.Println(err)
				continue
			}
			log.Println(string(content))
		}
	}
	return nil
}
