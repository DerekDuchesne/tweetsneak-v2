package db

import (
	"context"

	"cloud.google.com/go/datastore"

	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

var (
	client *datastore.Client
)

func init() {
	// Get the context.
	ctx := context.Background()

	// Create the datastore client.
	c, err := datastore.NewClient(ctx, settings.ProjectName)
	if err != nil {
		panic(err)
	}
	client = c
}
