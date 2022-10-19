package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"regexp"
	"time"
)

type Results struct {
	ResultCount int `json:"resultCount"`
	Results     []struct {
		WrapperType            string    `json:"wrapperType"`
		Kind                   string    `json:"kind"`
		CollectionID           int       `json:"collectionId"`
		TrackID                int       `json:"trackId"`
		ArtistName             string    `json:"artistName"`
		CollectionName         string    `json:"collectionName"`
		TrackName              string    `json:"trackName"`
		CollectionCensoredName string    `json:"collectionCensoredName"`
		TrackCensoredName      string    `json:"trackCensoredName"`
		CollectionViewURL      string    `json:"collectionViewUrl"`
		FeedURL                string    `json:"feedUrl"`
		TrackViewURL           string    `json:"trackViewUrl"`
		ArtworkURL30           string    `json:"artworkUrl30"`
		ArtworkURL60           string    `json:"artworkUrl60"`
		ArtworkURL100          string    `json:"artworkUrl100"`
		CollectionPrice        float64   `json:"collectionPrice"`
		TrackPrice             float64   `json:"trackPrice"`
		TrackRentalPrice       int       `json:"trackRentalPrice"`
		CollectionHdPrice      int       `json:"collectionHdPrice"`
		TrackHdPrice           int       `json:"trackHdPrice"`
		TrackHdRentalPrice     int       `json:"trackHdRentalPrice"`
		ReleaseDate            time.Time `json:"releaseDate"`
		CollectionExplicitness string    `json:"collectionExplicitness"`
		TrackExplicitness      string    `json:"trackExplicitness"`
		TrackCount             int       `json:"trackCount"`
		Country                string    `json:"country"`
		Currency               string    `json:"currency"`
		PrimaryGenreName       string    `json:"primaryGenreName"`
		ContentAdvisoryRating  string    `json:"contentAdvisoryRating"`
		ArtworkURL600          string    `json:"artworkUrl600"`
		GenreIds               []string  `json:"genreIds"`
		Genres                 []string  `json:"genres"`
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
		panic(err)
	}
	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	var response Results
	err = json.Unmarshal(body, &response)
	if err != nil {
		panic(err)
	}
	fmt.Println(response.Results[0].FeedURL)
}
