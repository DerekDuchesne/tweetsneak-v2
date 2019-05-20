package handlers

import (
	"fmt"
	"sort"
	"strings"
	"golang.org/x/sync/errgroup"

	"github.com/go-openapi/runtime/middleware"

	"github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/models"
	"github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/restapi/operations"
	"github.com/DerekDuchesne/tweetsneak-v2/dependencies/db"
	"github.com/DerekDuchesne/tweetsneak-v2/dependencies/logging"
	"github.com/DerekDuchesne/tweetsneak-v2/dependencies/tracing"
	"github.com/DerekDuchesne/tweetsneak-v2/dependencies/twitter"
	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

type getSearchHandler struct {
	Logger logging.Logger
	Tracer tracing.Tracer
	DatabaseClient db.DatabaseClient
	TwitterClient twitter.TwitterClient
}

type wordFrequency struct {
	word string
	count int
}

type wordFrequencies []wordFrequency

func (w wordFrequencies) Len() int { return len(w) }
func (w wordFrequencies) Less(i, j int) bool { return w[i].count > w[j].count }
func (w wordFrequencies) Swap(i, j int) { w[i], w[j] = w[j], w[i] }

func NewGetSearchHandler(logger logging.Logger, tracer tracing.Tracer, databaseClient db.DatabaseClient, twitterClient twitter.TwitterClient) *getSearchHandler {
	return &getSearchHandler{
		Logger: logger,
		Tracer: tracer,
		DatabaseClient: databaseClient,
		TwitterClient: twitterClient,
	}
}

func (g *getSearchHandler) Handle(params operations.GetSearchParams) middleware.Responder {
	// Get the request context.
	ctx := params.HTTPRequest.Context()
	
	// Get the keywords to search from the request.
	g.Tracer.StartSpan(ctx, "Read query parameter")
	q := params.Q
	g.Logger.Infof(ctx, "Read parameter q='%s' from the request URL.", q)
	g.Tracer.EndSpan(ctx)

	// Get the most recent tweets.
	var errGroup errgroup.Group
	tweetChan := make(chan twitter.Tweet)
	errGroup.Go(func() error {
		defer close(tweetChan)
		var maxID int64
		numTweets := 0

		for numTweets < settings.NumTweetsInResults {
			tweets, err := g.TwitterClient.Search(q, maxID)
			if err != nil {
				return err
			}
			if len(tweets) == 0 { // No more tweets to read.
				break
			}
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
	searchResult := new(models.SearchResult)
	wordCount := make(map[string]int)
	for tweet := range tweetChan {
		searchResult.Tweets = append(searchResult.Tweets, &models.Tweet{
			Text: tweet.Message,
			URL: fmt.Sprintf("https://twitter.com/i/web/status/%d", tweet.ID),
		})
		words := strings.Fields(stripChars(strings.ToLower(tweet.Message), "-=+|!@#$%^&*()`~[]{};:'\",<.>\\/?"))
		for _, word := range words {
			wordCount[word]++
		} 
	}

	// Check for errors while searching through tweets.
	if err := errGroup.Wait(); err != nil {
		g.Logger.Errorf(ctx, "Error searching tweets: %v", err)
		return operations.NewGetSearchInternalServerError().WithPayload(&models.Error{
			ErrorMessage: "Error searching for tweets.",
			ErrorType: "Server Error",
		})
	}

	// Get the most frequently occurring words.
	var frequencies wordFrequencies
	for word, count := range wordCount {
		frequencies = append(frequencies, wordFrequency{
			word: word,
			count: count,
		})
	}
	sort.Sort(frequencies)
	for i, frequency := range frequencies {
		if i == 10 {
			break
		}
		searchResult.Frequencies = append(searchResult.Frequencies, &models.WordFrequency{
			Count: int64(frequency.count),
			Word: frequency.word,
		})
	}

	return operations.NewGetSearchOK().WithPayload(searchResult)
}

func stripChars(str, chars string) string{
    return strings.Map(func(r rune) rune {
        if strings.IndexRune(chars, r) < 0 {
            return r
        }
        return -1
    }, str)
}