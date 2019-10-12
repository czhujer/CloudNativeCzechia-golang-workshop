package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// URL dwhere to read Atom Feed
var URL string

// entrypoint for  the program
func main() {
	log.Printf("Started parsing %s", URL)

	client := &http.Client{}

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
	fmt.Print(string(body))

}

// inicialization
func init() {
	flag.StringVar(&URL, "URL", "", "Atom feed url")
	flag.Parse()
}
