package gocqrs

import "reflect"

// Command represents any command object that can be handled by a CommandHandler.
// Commands are write operations that modify system state and may trigger side effects.
// Examples: CreateUserCommand, UpdateProductCommand, DeleteOrderCommand
type Command any

// CommandHandler defines the interface for handling commands.
// Each command type should have a corresponding handler that implements this interface.
// Handlers are responsible for executing business logic and collecting domain events.
type CommandHandler interface {
	// Handle processes the given command and returns the handler instance.
	// The handler pattern allows for method chaining and event collection.
	// Business logic should be executed within this method.
	Handle(c Command) CommandHandler

	// CollectEvents returns all domain events that were produced during command handling.
	// These events will be dispatched to the event bus after command execution.
	// Returns an empty slice if no events were produced.
	CollectEvents() []Event
}

// CommandBus defines the interface for a command bus that handles write operations.
// It provides methods to execute commands synchronously or asynchronously and register command handlers.
type CommandBus interface {
	// Dispatch executes a command asynchronously in a separate goroutine.
	// Use this for fire-and-forget operations or when you don't need to wait.
	// Any domain events produced will be dispatched to the event bus.
	Dispatch(c Command)

	// Execute executes a command synchronously and waits for completion.
	// Use this when you need to ensure the command has finished executing.
	// Any domain events produced will be dispatched to the event bus.
	Execute(c Command)

	// Register associates a command type with its corresponding handler.
	// The command parameter is used to determine the type name for registration.
	// Only one handler can be registered per command type (last registration wins).
	Register(c Command, ch CommandHandler)
}

// defaultCommandBus is the default implementation of CommandBus.
// It uses reflection to map command types to their handlers and integrates with an event bus.
type defaultCommandBus struct {
	// EventBus is used to dispatch domain events produced by command handlers
	EventBus EventBus
	// handlers maps command type names to their corresponding handlers
	handlers map[string]CommandHandler
}

// Dispatch executes the given command asynchronously in a new goroutine.
// It finds the registered handler, executes the command, and dispatches any resulting events.
// Note: This method returns immediately without waiting for command completion.
func (d *defaultCommandBus) Dispatch(c Command) {
	go d.handleCommand(c)
}

// Execute executes the given command synchronously.
// It finds the registered handler, executes the command, and dispatches any resulting events.
// Use this when you need to ensure the command has finished executing.
func (d *defaultCommandBus) Execute(c Command) {
	d.handleCommand(c)
}

// Register stores a command handler for the given command type.
// It uses reflection to extract the type name from the command instance.
// If a handler already exists for this command type, it will be replaced.
func (d *defaultCommandBus) Register(c Command, ch CommandHandler) {
	d.handlers[reflect.TypeOf(c).Name()] = ch
}

// handleCommand is the internal method that processes commands.
// It looks up the handler, executes the command, collects events, and dispatches them.
// Panics if no handler is registered for the command type.
func (d *defaultCommandBus) handleCommand(c Command) {
	ch := d.handlers[reflect.TypeOf(c).Name()]
	if ch == nil {
		panic("no handler registered for command type: " + reflect.TypeOf(c).Name())
	}
	ch = ch.Handle(c)
	events := ch.CollectEvents()
	for _, e := range events {
		d.EventBus.Dispatch(e)
	}
}

// DefaultCommandBus creates a new instance of the default command bus implementation.
// Requires an event bus instance for dispatching domain events produced by command handlers.
// Returns a CommandBus that uses reflection-based handler lookup.
func DefaultCommandBus(eventBus EventBus) *defaultCommandBus {
	return &defaultCommandBus{
		EventBus: eventBus,
		handlers: make(map[string]CommandHandler),
	}
}