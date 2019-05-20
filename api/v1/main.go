package main

import (
	"log"
	"os"

	"github.com/go-openapi/loads"

	"github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/restapi"
	"github.com/DerekDuchesne/tweetsneak-v2/api/v1/gen/restapi/operations"
	"github.com/DerekDuchesne/tweetsneak-v2/api/v1/handlers"
	"github.com/DerekDuchesne/tweetsneak-v2/dependencies/db"
	"github.com/DerekDuchesne/tweetsneak-v2/dependencies/logging"
	"github.com/DerekDuchesne/tweetsneak-v2/dependencies/tracing"
	"github.com/DerekDuchesne/tweetsneak-v2/dependencies/twitter"
)

func main() {
	// Load the API spec.
	swaggerSpec, err := loads.Analyzed(restapi.SwaggerJSON, "")
	if err != nil {
		log.Fatalln(err)
	}

	// Create the API.
	api := operations.NewSearchAPI(swaggerSpec)
	api.GetSearchHandler = handlers.NewGetSearchHandler(
		logging.NewStackDriverLogger(),
		tracing.NewStackDriverTracer(),
		db.NewGCDatastoreClient(),
		twitter.NewProductionTwitterClient(
			os.Getenv("TWITTER_CONSUMERKEY"),
			os.Getenv("TWITTER_CONSUMERSECRET"),
			os.Getenv("TWITTER_ACCESSTOKEN"),
			os.Getenv("TWITTER_ACCESSSECRET"),
		),
	)

	// Create the server.
	server := restapi.NewServer(api)
	defer server.Shutdown()

	// Start the server.
	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}
}