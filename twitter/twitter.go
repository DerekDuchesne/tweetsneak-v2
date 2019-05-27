package twitter

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"

	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

const (
	maxResults = 100
)

var (
	client *twitter.Client
)

// Tweet represents a single tweet from the Twitter API.
type Tweet struct {
	ID      int64
	Message string
}

func init() {
	// Create the Twitter API client.
	config := oauth1.NewConfig(settings.TwitterConsumerKey, settings.TwitterConsumerSecret)
	token := oauth1.NewToken(settings.TwitterAccessToken, settings.TwitterAccessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	client = twitter.NewClient(httpClient)
}

// Search will perform a single Twitter API search call with the given maxID as the starting point.
func Search(keyword string, maxID int64) ([]Tweet, error) {
	// Execute the search.
	search, resp, err := client.Search.Tweets(&twitter.SearchTweetParams{
		Query:      keyword,
		ResultType: "recent",
		Count:      maxResults,
		MaxID:      maxID,
	})
	if err != nil {
		return nil, err
	}

	// Handle errors.
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("Got status code %d. Response: %s", resp.StatusCode, string(body))
	}

	// Collect the tweets from the results.
	var tweets []Tweet
	for _, tweet := range search.Statuses {
		tweets = append(tweets, Tweet{
			ID:      tweet.ID,
			Message: tweet.Text,
		})
	}

	return tweets, nil
}
