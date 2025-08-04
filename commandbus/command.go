// Package commandbus provides a command bus implementation for CQRS (Command Query Responsibility Segregation) pattern.
// The command bus handles write operations that modify system state and may produce domain events.
package commandbus

import "github.com/avanboxel/gocqrs/eventbus"

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
	CollectEvents() []eventbus.Event
}
