package main

import (
	"os"
	"testing"
)

func TestGetRecentTweets(t *testing.T) {
	tweets, err := getRecentTweets(twitterCredentials{
		consumerKey:    os.Getenv("TWITTER_CONSUMERKEY"),
		consumerSecret: os.Getenv("TWITTER_CONSUMERSECRET"),
		accessToken:    os.Getenv("TWITTER_ACCESSTOKEN"),
		accessSecret:   os.Getenv("TWITTER_ACCESSSECRET"),
	})
	if err != nil {
		t.Fatal(err)
	}
	if len(tweets) != 5 {
		t.Fatalf("Expected 5 tweets to be returned. Got %d", len(tweets))
	}
}
