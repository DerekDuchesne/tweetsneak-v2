package db

import (
	"context"

	"cloud.google.com/go/datastore"

	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

type GCDatastoreClient struct {

}

func NewGCDatastoreClient() *GCDatastoreClient {
	return &GCDatastoreClient{}
}

func (g *GCDatastoreClient) Write(entityName string, obj interface{}) error {
	// Create the context.
	ctx := context.Background()

	// Create the datastore client.
	datastoreClient, err := datastore.NewClient(ctx, settings.ProjectName)
	if err != nil {
		return err
	}

	// Create the datastore key.
	key := datastore.IncompleteKey(entityName, nil)

	// Write the object to the database.
	if _, err := datastoreClient.Put(ctx, key, obj); err != nil {
		return err
	}

	return nil
}