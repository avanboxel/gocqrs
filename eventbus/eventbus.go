package eventbus

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

// NewDefault creates a new instance of the default event bus implementation.
// Returns an EventBus that uses string-based event type routing.
func NewDefault() *defaultEventBus {
	return &defaultEventBus{
		handlers: make(map[string]EventHandler),
	}
}
