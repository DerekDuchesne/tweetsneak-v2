package twitter

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

const (
	maxResults = 100
)

type ProductionTwitterClient struct {
	consumerKey string
	consumerSecret string
	accessToken string
	accessSecret string
}

func NewProductionTwitterClient(consumerKey, consumerSecret, accessToken, accessSecret string) *ProductionTwitterClient {
	return &ProductionTwitterClient{
		consumerKey: consumerKey,
		consumerSecret: consumerSecret,
		accessToken: accessToken,
		accessSecret: accessSecret,
	}
}

func (p *ProductionTwitterClient) Search(keyword string, maxID int64) ([]Tweet, error) {
	// Create the Twitter client.
	config := oauth1.NewConfig(p.consumerKey, p.consumerSecret)
	token := oauth1.NewToken(p.accessToken, p.accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient)

	// Execute the search.
	search, resp, err := twitterClient.Search.Tweets(&twitter.SearchTweetParams{
		Query: keyword,
		ResultType: "recent",
		Count: maxResults,
		MaxID: maxID,
	})
	if err != nil {
		return nil, err
	}
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
			ID: tweet.ID,
			Message: tweet.Text,
		})
	}

	return tweets, nil
}