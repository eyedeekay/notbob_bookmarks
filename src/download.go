package nbjson

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/eyedeekay/goSam"
)

func DownloadFile(url string, filename string) {
	if _, err := os.Stat(filename); !os.IsNotExist(err) {
		log.Printf("File %s already exists, skipping download", filename)
		return
	}

	sam, err := goSam.NewDefaultClient()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Client Created")

	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: sam.Dial,
		},
	}

	http.DefaultClient = httpClient
	log.Println("HTTP Client Created")

	resp, err := httpClient.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("File downloaded and saved as %s", filename)
}
