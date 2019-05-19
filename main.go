package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

type twitterCredentials struct {
	consumerKey    string
	consumerSecret string
	accessToken    string
	accessSecret   string
}

func main() {
	// Define the credentials.
	creds := twitterCredentials{
		consumerKey:    os.Getenv("TWITTER_CONSUMERKEY"),
		consumerSecret: os.Getenv("TWITTER_CONSUMERSECRET"),
		accessToken:    os.Getenv("TWITTER_ACCESSTOKEN"),
		accessSecret:   os.Getenv("TWITTER_ACCESSSECRET"),
	}

	// Search the tweets.
	tweets, err := getRecentTweets(creds)
	if err != nil {
		log.Fatalf("Error searching recent tweets: %v", err)
	}

	fmt.Printf("%#v\n", tweets)
}

func getRecentTweets(creds twitterCredentials) ([]string, error) {
	// Create the Twitter client.
	config := oauth1.NewConfig(creds.consumerKey, creds.consumerSecret)
	token := oauth1.NewToken(creds.accessToken, creds.accessSecret)
	httpClient := config.Client(oauth1.NoContext, token)
	twitterClient := twitter.NewClient(httpClient)

	// Search the tweets.
	var tweets []string
	var maxID int64
	for len(tweets) < 5 {
		search, _, err := twitterClient.Search.Tweets(&twitter.SearchTweetParams{
			Query:      "test",
			ResultType: "recent",
			Count:      100,
			MaxID:      maxID,
		})
		if err != nil {
			return nil, err
		}
		for _, tweet := range search.Statuses {
			tweets = append(tweets, tweet.Text)
			maxID = tweet.ID - 1
			if len(tweets) == 5 {
				break
			}
		}
	}

	return tweets, nil
}
