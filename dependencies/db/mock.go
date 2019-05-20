package db

import (
	"fmt"
)

type MockDatabaseClient struct {
	shouldError bool
}

func NewMockDatabaseClient(shouldError bool) *MockDatabaseClient {
	return &MockDatabaseClient{
		shouldError: shouldError,
	}
}

func (m *MockDatabaseClient) Write(entityName string, obj interface{}) error {
	if m.shouldError {
		return fmt.Errorf("Generic database error")
	}
	return nil
}