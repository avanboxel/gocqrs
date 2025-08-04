package main

import (
	"fmt"

	"github.com/avanboxel/gocqrs/querybus"
)

type GetUsernameQuery struct {
	ID int
}

type GetUsernameQueryHandler struct{}

func (h *GetUsernameQueryHandler) Handle(q querybus.Query) querybus.QueryResult {
	getUsernameQuery := q.(GetUsernameQuery)

	// Fake username data based on ID
	usernames := map[int]string{
		1: "john_doe",
		2: "jane_smith",
		3: "bob_wilson",
		4: "alice_johnson",
		5: "charlie_brown",
	}

	if username, exists := usernames[getUsernameQuery.ID]; exists {
		return querybus.QueryResult{
			Payload: username,
			Success: true,
		}
	}

	return querybus.QueryResult{
		Payload: "unknown_user",
		Success: false,
	}
}

func main() {
	queryBus := querybus.NewDefault()

	// Register the query handler
	queryBus.Register(GetUsernameQuery{}, &GetUsernameQueryHandler{})

	// Ask for usernames
	for i := 1; i <= 6; i++ {
		query := GetUsernameQuery{ID: i}
		result := queryBus.Ask(query)
		if result.Success {
			fmt.Printf("User ID %d: %s (found)\n", i, result.Payload)
		} else {
			fmt.Printf("User ID %d: %s (not found)\n", i, result.Payload)
		}
	}
}
