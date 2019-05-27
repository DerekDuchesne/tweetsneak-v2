package settings

import (
	"os"
	"strconv"
	"strings"
)

var (
	// ENVIRONMENT VARIABLE NAMES

	// ProjectNameEnvName is the environment variable name for ProjectName.
	ProjectNameEnvName = "GAE_PROJECT_NAME"
	// UseProductionEnvName is the environment variable name for UseProduction.
	UseProductionEnvName = "USE_PRODUCTION"
	// NumTweetsInResultsEnvName is the environment variable name for NumTweetsInResults.
	NumTweetsInResultsEnvName = "NUM_TWEETS_IN_RESULTS"
	// NumWordFrequenciesInResultsEnvName is the environment variable name for NumWordFrequenciesInResults.
	NumWordFrequenciesInResultsEnvName = "NUM_WORD_FREQUENCIES_IN_RESULTS"
	// TwitterConsumerKeyEnvName is the environment variable name for TwitterConsumerKey.
	TwitterConsumerKeyEnvName = "TWITTER_CONSUMERKEY"
	// TwitterConsumerSecretEnvName is the environment variable name for TwitterConsumerSecret.
	TwitterConsumerSecretEnvName = "TWITTER_CONSUMERSECRET"
	// TwitterAccessTokenEnvName is the environment variable name for TwitterAccessToken.
	TwitterAccessTokenEnvName = "TWITTER_ACCESSTOKEN"
	// TwitterAccessSecretEnvName is the environment variable name for TwitterAccessSecret.
	TwitterAccessSecretEnvName = "TWITTER_ACCESSSECRET"

	// ENVIRONMENT VARIABLE VALUES

	// ProjectName is the name of the App Engine project.
	ProjectName = "tweetsneak-v2"
	// UseProduction is a boolean flag determining whether or not production services should be used.
	UseProduction = false
	// NumTweetsInResults is the number of tweets to display on the web page and use when searching for the words with the highest frequency.
	NumTweetsInResults = 1000
	// NumWordFrequenciesInResults is the number of most frequent words we should display in the results.
	NumWordFrequenciesInResults = 10
	// TwitterConsumerKey is part of our Twitter API credentials.
	TwitterConsumerKey = ""
	// TwitterConsumerSecret is part of our Twitter API credentials.
	TwitterConsumerSecret = ""
	// TwitterAccessToken is part of our Twitter API credentials.
	TwitterAccessToken = ""
	// TwitterAccessSecret is part of our Twitter API credentials.
	TwitterAccessSecret = ""
)

func init() {
	// Get settings from environment variables
	projectName, ok := os.LookupEnv(ProjectNameEnvName)
	if ok {
		ProjectName = projectName
	}

	useProduction, ok := os.LookupEnv(UseProductionEnvName)
	if ok {
		UseProduction = strings.ToLower(useProduction) == "true"
	}

	numTweetsInResults, ok := os.LookupEnv(NumTweetsInResultsEnvName)
	if ok {
		parsedNumTweetsInResults, err := strconv.Atoi(numTweetsInResults)
		if err == nil {
			NumTweetsInResults = parsedNumTweetsInResults
		}
	}

	numWordFrequenciesInResults, ok := os.LookupEnv(NumWordFrequenciesInResultsEnvName)
	if ok {
		parsedNumFrequenciesInResults, err := strconv.Atoi(numWordFrequenciesInResults)
		if err == nil {
			NumWordFrequenciesInResults = parsedNumFrequenciesInResults
		}
	}

	twitterConsumerKey, ok := os.LookupEnv(TwitterConsumerKeyEnvName)
	if ok {
		TwitterConsumerKey = twitterConsumerKey
	}

	twitterConsumerSecret, ok := os.LookupEnv(TwitterConsumerSecretEnvName)
	if ok {
		TwitterConsumerSecret = twitterConsumerSecret
	}

	twitterAccessToken, ok := os.LookupEnv(TwitterAccessTokenEnvName)
	if ok {
		TwitterAccessToken = twitterAccessToken
	}

	twitterAccessSecret, ok := os.LookupEnv(TwitterAccessSecretEnvName)
	if ok {
		TwitterAccessSecret = twitterAccessSecret
	}
}
