package gocqrs

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

// EventBus defines the interface for an event bus that handles domain event dispatch.
// It provides methods to dispatch events to registered handlers and register event handlers.
type EventBus interface {
	// Dispatch sends an event to its registered handler.
	// The event's GetEventType() method is used to find the appropriate handler.
	// Panics if no handler is registered for the event type.
	Dispatch(e Event)
	
	// Register associates an event type with its corresponding handler.
	// The eventType should match the string returned by Event.GetEventType().
	// Only one handler can be registered per event type (last registration wins).
	Register(eventType string, eh EventHandler)
}

// defaultEventBus is the default implementation of EventBus.
// It uses a simple map to route events to their handlers based on event type.
type defaultEventBus struct {
	// handlers maps event type strings to their corresponding handlers
	handlers map[string]EventHandler
}

// Dispatch sends the given event to its registered handler.
// Uses the event's GetEventType() method to look up the handler.
// Panics if no handler is registered for the event type.
func (d *defaultEventBus) Dispatch(e Event) {
	handler := d.handlers[e.GetEventType()]
	if handler == nil {
		panic("no handler registered for event type: " + e.GetEventType())
	}
	handler(e)
}

// Register stores an event handler for the given event type.
// The eventType parameter should match what Event.GetEventType() returns.
// If a handler already exists for this event type, it will be replaced.
func (d *defaultEventBus) Register(eventType string, eh EventHandler) {
	d.handlers[eventType] = eh
}

// DefaultEventBus creates a new instance of the default event bus implementation.
// Returns an EventBus that uses string-based event type routing.
func DefaultEventBus() *defaultEventBus {
	return &defaultEventBus{
		handlers: make(map[string]EventHandler),
	}
}