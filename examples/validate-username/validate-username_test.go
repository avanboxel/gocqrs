package validateusername

import (
	"testing"

	"github.com/avanboxel/gocqrs/commandbus"
	"github.com/avanboxel/gocqrs/eventbus"
)

func TestValidateUsernameValid(t *testing.T) {
	// Track events received
	var receivedEvents []eventbus.Event

	// Create event bus and register event handler
	eventBus := eventbus.NewDefault()
	eventBus.Register("UsernameValidated", func(e eventbus.Event) {
		receivedEvents = append(receivedEvents, e)
	})

	// Create command bus
	commandBus := commandbus.NewDefault(eventBus)

	// Create valid username command (8-16 characters)
	validateCmd := ValidateUsernameCommand{
		Username: "validuser123",
	}

	validateHandler := ValidateUsernameHandler{
		events: []eventbus.Event{},
	}

	// Register the command with its handler
	commandBus.Register(validateCmd, validateHandler)

	// Dispatch the command
	commandBus.Dispatch(validateCmd)

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
	var receivedEvents []eventbus.Event

	// Create event bus and register event handler
	eventBus := eventbus.NewDefault()
	eventBus.Register("UsernameValidated", func(e eventbus.Event) {
		receivedEvents = append(receivedEvents, e)
	})

	// Create command bus
	commandBus := commandbus.NewDefault(eventBus)

	// Create short username command (less than 8 characters)
	validateCmd := ValidateUsernameCommand{
		Username: "short",
	}

	validateHandler := ValidateUsernameHandler{
		events: []eventbus.Event{},
	}

	// Register the command with its handler
	commandBus.Register(validateCmd, validateHandler)

	// Dispatch the command
	commandBus.Dispatch(validateCmd)

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
	var receivedEvents []eventbus.Event

	// Create event bus and register event handler
	eventBus := eventbus.NewDefault()
	eventBus.Register("UsernameValidated", func(e eventbus.Event) {
		receivedEvents = append(receivedEvents, e)
	})

	// Create command bus
	commandBus := commandbus.NewDefault(eventBus)

	// Create long username command (more than 16 characters)
	validateCmd := ValidateUsernameCommand{
		Username: "verylongusernamethatistoolong",
	}

	validateHandler := ValidateUsernameHandler{
		events: []eventbus.Event{},
	}

	// Register the command with its handler
	commandBus.Register(validateCmd, validateHandler)

	// Dispatch the command
	commandBus.Dispatch(validateCmd)

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
