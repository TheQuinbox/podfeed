package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"regexp"
)

type Results struct {
	Results []struct {
		FeedURL string `json:"feedUrl"`
	} `json:"results"`
}

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <URL>", os.Args[0])
		os.Exit(1)
	}
	re := regexp.MustCompile(`[^w]+\/id(?P<id>\d+)`)
	matches := re.FindStringSubmatch(os.Args[1])
	id := matches[1]
	url := fmt.Sprintf("https://itunes.apple.com/lookup?id=%s&entity=podcast", id)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	var response Results
	err = json.Unmarshal(body, &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Results[0].FeedURL)
}
