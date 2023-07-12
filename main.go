package main

import (
	"log"

	nbjson "github.com/eyedeekay/notbob_bookmarks.json/src"
)

func main() {
	nbjson.DownloadFile("http://notbob.i2p/blog.xml", "notbob.xml")
	xmlfile, err := nbjson.LoadXMLFile("notbob.xml")
	if err != nil {
		log.Fatal(err)
	}
	bookmarks := nbjson.LoopOverXMLFile(xmlfile)
	jsonBookmarks := nbjson.BookmarksToJSONFeed("Not Bob's Bookmarks", bookmarks)
	bytes := nbjson.PrettyPrintJSONFeed(jsonBookmarks)
	nbjson.SaveJSONFeed(bytes)
	allfeeds := nbjson.CombineAllFeeds()
	log.Println(string(allfeeds))
}
