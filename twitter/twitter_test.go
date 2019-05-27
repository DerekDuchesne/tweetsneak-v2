package twitter

import "testing"

func TestSearch(t *testing.T) {
	result, err := Search("test", 0)
	if err != nil {
		t.Fatalf("Got error when trying to execute a search: %v", err)
	}

	if len(result) == 0 {
		t.Fatalf("Got 0 search results from Twitter API when there should have been at least one.")
	}
}
