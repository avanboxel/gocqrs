// Package eventbus provides an event bus implementation for domain event handling.
// The event bus enables decoupled communication between different parts of the system
// through the publication and handling of domain events.
package eventbus

// Event represents a domain event that has occurred in the system.
// Events are immutable facts about something that has happened.
// Examples: UserCreatedEvent, OrderProcessedEvent, PaymentFailedEvent
type Event interface {
	// GetEventType returns a string identifier for the event type.
	// This is used by the event bus to route events to the appropriate handlers.
	// Should return a consistent string for each event type (e.g., "UserCreated").
	GetEventType() string
}

// EventHandler defines a function type for handling domain events.
// Event handlers should be side-effect free and idempotent when possible.
// They are called synchronously when events are dispatched.
type EventHandler func(e Event)
