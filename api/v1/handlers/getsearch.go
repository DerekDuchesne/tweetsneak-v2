package handlers

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/go-openapi/runtime/middleware"
	"go.opencensus.io/trace"
	"golang.org/x/sync/errgroup"

	apimodels "github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/models"
	"github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/restapi/operations"
	"github.com/DerekDuchesne/tweetsneak-v2/db"
	"github.com/DerekDuchesne/tweetsneak-v2/logging"
	"github.com/DerekDuchesne/tweetsneak-v2/settings"
	"github.com/DerekDuchesne/tweetsneak-v2/twitter"
	"github.com/DerekDuchesne/tweetsneak-v2/utils"
)

type wordFrequency struct {
	word  string
	count int
}

type wordFrequencies []wordFrequency

func (w wordFrequencies) Len() int           { return len(w) }
func (w wordFrequencies) Less(i, j int) bool { return w[i].count > w[j].count }
func (w wordFrequencies) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }

// Search will read through tweets from the Twitter API, count the most frequent words, and return the results.
func Search(params operations.GetSearchParams) middleware.Responder {
	// Get the request context.
	ctx := params.HTTPRequest.Context()

	// Get the keywords to search from the request.
	_, span := trace.StartSpan(ctx, "Read query parameter")
	q := params.Q
	span.End()
	logging.Infof(ctx, "Read parameter q='%s' from the request URL.", q)

	// Get the most recent tweets.
	var errGroup errgroup.Group
	tweetChan := make(chan twitter.Tweet)
	errGroup.Go(func() error {
		// Close the channel once this goroutine complete.
		defer close(tweetChan)

		// Create a span for the entire goroutine.
		ctx, span := trace.StartSpan(ctx, "Search tweets")
		defer span.End()

		// Keep track of and limit the number of API requests that will be made.
		var maxID int64
		numTweets := 0

		for numTweets < settings.NumTweetsInResults {
			// Make the API call.
			tweets, more, err := searchTweets(ctx, q, maxID)
			if err != nil {
				return err
			}

			// Stop if there are no more tweets to search through.
			if !more {
				break
			}

			// Pass each tweet through the channel.
			for _, tweet := range tweets {
				maxID = tweet.ID
				tweetChan <- tweet
				numTweets++
				if numTweets == settings.NumTweetsInResults {
					break
				}
			}
		}

		return nil
	})

	// Collect all of the tweets and count the words.
	_, span = trace.StartSpan(ctx, "Count word frequencies")
	var tweets []twitter.Tweet
	wordCount := make(map[string]int)
	for tweet := range tweetChan {
		tweets = append(tweets, tweet)
		words := strings.Fields(utils.StripChars(strings.ToLower(tweet.Message), "-=+|!@#$%^&*()`~[]{};:'\",<.>\\/?"))
		for _, word := range words {
			wordCount[word]++
		}
	}
	span.End()

	// Check for errors while searching through tweets.
	if err := errGroup.Wait(); err != nil {
		logging.Errorf(ctx, "Error searching tweets: %v", err)
		return operations.NewGetSearchInternalServerError().WithPayload(&apimodels.Error{
			ErrorMessage: "Error searching for tweets.",
			ErrorType:    "Server Error",
		})
	}

	// Get the most frequently occurring words.
	_, span = trace.StartSpan(ctx, "Get most frequent words.")
	var frequencies wordFrequencies
	for word, count := range wordCount {
		frequencies = append(frequencies, wordFrequency{
			word:  word,
			count: count,
		})
	}
	sort.Sort(frequencies)
	frequencies = frequencies[0:utils.Min(settings.NumWordFrequenciesInResults, len(frequencies))]
	span.End()
	logging.Infof(ctx, "Found %d tweets.", len(tweets))

	// Write the search result to the database.
	_, span = trace.StartSpan(ctx, "Write search result to database")
	searchEntity := new(db.SearchResult)
	searchEntity.Keyword = q
	for _, frequency := range frequencies {
		searchEntity.Frequencies = append(searchEntity.Frequencies, db.WordFrequency{
			Word:  frequency.word,
			Count: frequency.count,
		})
	}
	if err := searchEntity.Create(ctx); err != nil {
		logging.Errorf(ctx, "Error storing search result: %v", err)
	} else {
		logging.Infof(ctx, "Wrote results to database.")
	}
	span.End()

	// Write the output.
	searchResult := new(apimodels.SearchResult)
	for _, tweet := range tweets {
		searchResult.Tweets = append(searchResult.Tweets, &apimodels.Tweet{
			Text: tweet.Message,
			URL:  fmt.Sprintf("https://twitter.com/i/web/status/%d", tweet.ID),
		})
	}
	for _, frequency := range frequencies {
		searchResult.Frequencies = append(searchResult.Frequencies, &apimodels.WordFrequency{
			Word:  frequency.word,
			Count: int64(frequency.count),
		})
	}

	return operations.NewGetSearchOK().WithPayload(searchResult)
}

func searchTweets(ctx context.Context, q string, maxID int64) ([]twitter.Tweet, bool, error) {
	// Start the span for the API call.
	_, span := trace.StartSpan(ctx, "Twitter API call")
	defer span.End()

	tweets, err := twitter.Search(q, maxID)
	if err != nil {
		return nil, false, err
	}

	if len(tweets) == 0 {
		return nil, false, nil
	}

	return tweets, true, nil
}
