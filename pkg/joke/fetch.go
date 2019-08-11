package joke

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gabrie30/joke/pkg/date"
)

// FetchIfNeeded fetches jokes once per day
func FetchIfNeeded() {
	if needToFetch() == false {
		return
	}

	n := saveFilteredJokes()
	err := date.NewFetch(n)
	if err != nil {
		// TODO, if there is an error, delete last n jokes saved so there are no duplicate jokes
	}
}

func saveFilteredJokes() int {
	jokes := fetchAndFilterJokes()
	jokesSaved := 0
	for _, j := range jokes {

		err := j.Save()

		if err != nil {
			// TODO only log errors that are not uniqueness constraints
			continue
		}

		jokesSaved++
	}

	return jokesSaved
}

func needToFetch() bool {
	todaysDate := time.Now().Format("01-02-2006")
	if date.LastFetchDate() == todaysDate {
		return false
	}

	return true
}

func fetchAndFilterJokes() []Data {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://reddit.com/r/Jokes.json", nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Golang_Spider_Bot/3.0")

	resp, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Fatalln(err)
	}

	var d payload

	err = json.Unmarshal(body, &d)

	if err != nil {
		log.Fatal(err)
	}

	filteredJokes := []Data{}

	for _, j := range d.RedditDataPayload.RedditChild {

		if len(j.RedditData.Setup) < 110 && len(j.RedditData.Punchline) < 200 {
			if strings.Contains(j.RedditData.Punchline, "https://discord.gg/jokes") {
				continue
			}

			if strings.ContainsAny(j.RedditData.Setup, charsNotAllowed) || strings.ContainsAny(j.RedditData.Punchline, charsNotAllowed) {
				continue
			}

			if j.RedditData.WhitelistStatus != "all_ads" {
				continue
			}

			if j.RedditData.AdultsOnly == true {
				continue
			}
			filteredJokes = append(filteredJokes, j.RedditData)
		}

	}

	return filteredJokes
}
