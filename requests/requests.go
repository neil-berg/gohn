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

// GetTopStories fetches and formats a specified count of top Hacker News
// stories
func GetTopStories(count *int) {
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
	t := time.Unix(int64(story.Time), 0)
	tFmt := fmt.Sprintf("%s %02d, %04d %02d:%02d UTC",
		t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
	fmt.Printf("#%d \t%s\n  \tScore:\t %d\n  \tPosted:\t %s\n  \t%s\n\n",
		i+1, story.Title, story.Score, tFmt, story.URL)
}
