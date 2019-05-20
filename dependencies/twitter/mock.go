package twitter

import (
	"fmt"
)

type MockTwitterClient struct {
	shouldError bool
}

func NewMockTwitterClient(shouldError bool) *MockTwitterClient {
	return &MockTwitterClient{
		shouldError: shouldError,
	}
}

func (m *MockTwitterClient) Search(keyword string, maxID int64) ([]Tweet, error) {
	if m.shouldError {
		return nil, fmt.Errorf("Generic Twitter error")
	}
	return []Tweet{
		Tweet{
			ID: 0,
			Message: "This is a test.",
		},
	}, nil
}