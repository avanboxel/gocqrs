package userregister

import (
	"testing"

	"github.com/avanboxel/gocqrs/commandbus"
	"github.com/avanboxel/gocqrs/eventbus"
)

func TestDispatchUserRegister(t *testing.T) {
	// Track events received
	var receivedEvents []eventbus.Event

	// Create event bus and register event handler
	eventBus := eventbus.NewDefault()
	eventBus.Register("UserRegistered", func(e eventbus.Event) {
		receivedEvents = append(receivedEvents, e)
	})

	// Create command bus
	commandBus := commandbus.NewDefault(eventBus)

	// Create command and handler
	registerCmd := RegisterCommand{
		Username: "testuser",
		Email:    "test@example.com",
		Password: "password123",
	}

	registerHandler := RegisterCommandHandler{
		events: []eventbus.Event{},
	}

	// Register the command with its handler
	commandBus.Register(registerCmd, registerHandler)

	// Dispatch the command
	commandBus.Dispatch(registerCmd)

	// Verify event was dispatched
	if len(receivedEvents) != 1 {
		t.Errorf("Expected 1 event, got %d", len(receivedEvents))
	}

	// Verify the event is UserRegistered
	if userRegistered, ok := receivedEvents[0].(UserRegistered); ok {
		if userRegistered.Username != "testuser" {
			t.Errorf("Expected username 'testuser', got '%s'", userRegistered.Username)
		}
		if userRegistered.Email != "test@example.com" {
			t.Errorf("Expected email 'test@example.com', got '%s'", userRegistered.Email)
		}
	} else {
		t.Error("Expected UserRegistered event")
	}
}
