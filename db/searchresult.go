package db

import (
	"context"
	"log"
	"time"

	"cloud.google.com/go/datastore"

	"github.com/DerekDuchesne/tweetsneak-v2/settings"
)

// SearchResultKind is the name of the SearchResult model in datastore.
const SearchResultKind = "SearchResult"

// SearchResult is an entity in datastore that stores searches.
type SearchResult struct {
	Keyword     string          `datastore:"k"`
	Frequencies []WordFrequency `datastore:"f"`
	DateCreated time.Time       `datastore:"dc"`
}

// WordFrequency is an object containing a word and how often that word appears in the search results.
type WordFrequency struct {
	Word  string `datastore:"w"`
	Count int    `datastore:"c"`
}

// Create will put the given search result in datastore.
func (s *SearchResult) Create(ctx context.Context) error {
	// Create the datastore key.
	key := datastore.IncompleteKey(SearchResultKind, nil)

	// Set the DateCreated field to right now.
	s.DateCreated = time.Now()

	// If we're not on production, just log that we wrote the entity.
	if !settings.UseProduction {
		log.Printf("Wrote Search Result %#v to database.\n", s)
		return nil
	}

	// Put the SearchResult in datastore.
	_, err := client.Put(ctx, key, s)
	return err
}
