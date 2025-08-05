package validateusername

import (
	"testing"

	"github.com/avanboxel/gocqrs"
)

func TestValidateUsernameValid(t *testing.T) {
	// Track events received
	var receivedEvents []gocqrs.Event

	// Create event bus and register event handler
	eventBus := gocqrs.DefaultSyncEventBus()
	eventBus.Register("UsernameValidated", func(e gocqrs.Event) {
		receivedEvents = append(receivedEvents, e)
	})

	// Create command bus
	commandBus := gocqrs.DefaultCommandBus(eventBus)

	// Create valid username command (8-16 characters)
	validateCmd := ValidateUsernameCommand{
		Username: "validuser123",
	}

	validateHandler := &ValidateUsernameHandler{
		events: []gocqrs.Event{},
	}

	// Register the command with its handler
	commandBus.Register(validateCmd, validateHandler)

	// Execute the command synchronously
	commandBus.Execute(validateCmd)

	// Verify event was dispatched
	if len(receivedEvents) != 1 {
		t.Errorf("Expected 1 event, got %d", len(receivedEvents))
	}

	// Verify the event is UsernameValidated and valid
	if validated, ok := receivedEvents[0].(UsernameValidated); ok {
		if validated.Username != "validuser123" {
			t.Errorf("Expected username 'validuser123', got '%s'", validated.Username)
		}
		if !validated.IsValid {
			t.Errorf("Expected username to be valid, but got invalid with reason: %s", validated.Reason)
		}
		if validated.Reason != "Username is valid" {
			t.Errorf("Expected reason 'Username is valid', got '%s'", validated.Reason)
		}
	} else {
		t.Error("Expected UsernameValidated event")
	}
}

func TestValidateUsernameTooShort(t *testing.T) {
	// Track events received
	var receivedEvents []gocqrs.Event

	// Create event bus and register event handler
	eventBus := gocqrs.DefaultSyncEventBus()
	eventBus.Register("UsernameValidated", func(e gocqrs.Event) {
		receivedEvents = append(receivedEvents, e)
	})

	// Create command bus
	commandBus := gocqrs.DefaultCommandBus(eventBus)

	// Create short username command (less than 8 characters)
	validateCmd := ValidateUsernameCommand{
		Username: "short",
	}

	validateHandler := &ValidateUsernameHandler{
		events: []gocqrs.Event{},
	}

	// Register the command with its handler
	commandBus.Register(validateCmd, validateHandler)

	// Execute the command synchronously
	commandBus.Execute(validateCmd)

	// Verify event was dispatched
	if len(receivedEvents) != 1 {
		t.Errorf("Expected 1 event, got %d", len(receivedEvents))
	}

	// Verify the event is UsernameValidated and invalid
	if validated, ok := receivedEvents[0].(UsernameValidated); ok {
		if validated.Username != "short" {
			t.Errorf("Expected username 'short', got '%s'", validated.Username)
		}
		if validated.IsValid {
			t.Error("Expected username to be invalid")
		}
		if validated.Reason != "Username too short (minimum 8 characters)" {
			t.Errorf("Expected reason 'Username too short (minimum 8 characters)', got '%s'", validated.Reason)
		}
	} else {
		t.Error("Expected UsernameValidated event")
	}
}

func TestValidateUsernameTooLong(t *testing.T) {
	// Track events received
	var receivedEvents []gocqrs.Event

	// Create event bus and register event handler
	eventBus := gocqrs.DefaultSyncEventBus()
	eventBus.Register("UsernameValidated", func(e gocqrs.Event) {
		receivedEvents = append(receivedEvents, e)
	})

	// Create command bus
	commandBus := gocqrs.DefaultCommandBus(eventBus)

	// Create long username command (more than 16 characters)
	validateCmd := ValidateUsernameCommand{
		Username: "verylongusernamethatistoolong",
	}

	validateHandler := &ValidateUsernameHandler{
		events: []gocqrs.Event{},
	}

	// Register the command with its handler
	commandBus.Register(validateCmd, validateHandler)

	// Execute the command synchronously
	commandBus.Execute(validateCmd)

	// Verify event was dispatched
	if len(receivedEvents) != 1 {
		t.Errorf("Expected 1 event, got %d", len(receivedEvents))
	}

	// Verify the event is UsernameValidated and invalid
	if validated, ok := receivedEvents[0].(UsernameValidated); ok {
		if validated.Username != "verylongusernamethatistoolong" {
			t.Errorf("Expected username 'verylongusernamethatistoolong', got '%s'", validated.Username)
		}
		if validated.IsValid {
			t.Error("Expected username to be invalid")
		}
		if validated.Reason != "Username too long (maximum 16 characters)" {
			t.Errorf("Expected reason 'Username too long (maximum 16 characters)', got '%s'", validated.Reason)
		}
	} else {
		t.Error("Expected UsernameValidated event")
	}
}
