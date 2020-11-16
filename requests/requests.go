package requests

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// Story JSON structure from the API call
type Story struct {
	By          string `json:"by"`
	Descendants int    `json:"descendants"`
	ID          int    `json:"id"`
	Kids        []int  `json:"kids"`
	Score       int    `json:"score"`
	Time        int    `json:"time"`
	Title       string `json:"title"`
	Type        string `json:"type"`
	URL         string `json:"url"`
}

// GetTopStoryIDs fetches and formats a specified count of top Hacker News
// stories
func GetTopStoryIDs(count *int) {
	client := http.Client{Timeout: time.Second * 2}
	url := "https://hacker-news.firebaseio.com/v0/topstories.json"
	res, err := client.Get(url)
	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	var topStoryIDs []int
	decoder := json.NewDecoder(res.Body)
	decodeErr := decoder.Decode(&topStoryIDs)

	if decodeErr != nil {
		log.Fatal(decodeErr)
	}

	getStories(topStoryIDs[:*count])
}

func getStories(IDs []int) {
	client := http.Client{Timeout: time.Second * 2}
	baseURL := "https://hacker-news.firebaseio.com/v0/item/"

	for i := 0; i < len(IDs); i++ {
		url := baseURL + strconv.Itoa(IDs[i]) + ".json"
		res, err := client.Get(url)
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		body, bodyErr := ioutil.ReadAll(res.Body)
		if bodyErr != nil {
			log.Fatal(bodyErr)
		}

		var story Story
		jsonErr := json.Unmarshal(body, &story)
		if jsonErr != nil {
			log.Fatal(jsonErr)
		}

		formatStory(i, story)
	}

}

func formatStory(i int, story Story) {
	// TODO: Format the posted time nicely
	fmt.Printf("#%d \t%s\n  \tScore:\t %d\n  \tPosted:\t %d\n  \t%s\n\n",
		i+1, story.Title, story.Score, story.Time, story.URL)
}
