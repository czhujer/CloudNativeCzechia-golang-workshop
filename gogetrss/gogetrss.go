package main

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// URL dwhere to read Atom Feed
var URL string

type Feed struct {
	Entries []FeedEntry `xml:"entry"`
}

type FeedEntry struct {
	Title   string   `xml:"title"`
	Updated string   `xml:"updated"`
	Link    LinkItem `xml:"link,attr"`
}

type LinkItem struct {
	Link string `xml:"href,attr"`
}

// entrypoint for  the program
func main() {
	log.Printf("Started parsing %s", URL)

	// https://golang.org/pkg/net/http/#pkg-overview
	client := http.Client{}

	req, err := http.NewRequest("GET", URL, nil)
	if err != nil {
		log.Fatalf("ERROR: Cannot create request %s", URL)
	}

	req.Header.Add("User-Agent", "Golang Bot 1.0")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalf("Cannot get response %s", err.Error())
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Cannot read response %s", err.Error())
	}

	//fmt.Print(string(body))

	// xml decode
	// https://golang.org/pkg/encoding/xml/
	feed := Feed{}
	err = xml.Unmarshal(body, &feed)
	if err != nil {
		log.Fatalf("Cannot parse response: %s", err.Error())
	}

	for index, item := range feed.Entries {
		fmt.Printf("%d: %s, %s, %s\n", index, item.Updated, item.Title, item.Link.Link)
	}

}

// inicialization
func init() {
	flag.StringVar(&URL, "URL", "", "Atom feed url")
	flag.Parse()
}
