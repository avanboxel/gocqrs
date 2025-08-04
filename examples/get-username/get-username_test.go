package main

import (
	"testing"

	"github.com/avanboxel/gocqrs/querybus"
)

func TestGetUsernameQuery(t *testing.T) {
	queryBus := querybus.NewDefault()
	queryBus.Register(GetUsernameQuery{}, &GetUsernameQueryHandler{})

	tests := []struct {
		id              int
		expectedPayload string
		expectedSuccess bool
	}{
		{1, "john_doe", true},
		{2, "jane_smith", true},
		{3, "bob_wilson", true},
		{4, "alice_johnson", true},
		{5, "charlie_brown", true},
		{999, "unknown_user", false}, // Non-existent user
	}

	for _, test := range tests {
		query := GetUsernameQuery{ID: test.id}
		result := queryBus.Ask(query)

		if result.Success != test.expectedSuccess {
			t.Errorf("For ID %d, expected success %v, got %v", test.id, test.expectedSuccess, result.Success)
			continue
		}

		username, ok := result.Payload.(string)
		if !ok {
			t.Errorf("Expected string payload for ID %d, got %T", test.id, result.Payload)
			continue
		}

		if username != test.expectedPayload {
			t.Errorf("For ID %d, expected %s, got %s", test.id, test.expectedPayload, username)
		}
	}
}
