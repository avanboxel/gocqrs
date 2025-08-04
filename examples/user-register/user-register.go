package userregister

import (
	"fmt"

	"github.com/avanboxel/gocqrs/commandbus"
	"github.com/avanboxel/gocqrs/eventbus"
)

type RegisterCommand struct {
	Username string
	Email    string
	Password string
}

type UserRegistered struct {
	Username string
	Email    string
}

func (e UserRegistered) GetEventType() string {
	return "UserRegistered"
}

type RegisterCommandHandler struct {
	events []eventbus.Event
}

func (r RegisterCommandHandler) Handle(c commandbus.Command) commandbus.CommandHandler {
	if cmd, ok := c.(RegisterCommand); ok {
		// Fake user registration logic
		fmt.Printf("Registering user: %s with email: %s\n", cmd.Username, cmd.Email)
		fmt.Println("Save user here")

		// Create and add the UserRegistered event
		userRegisteredEvent := UserRegistered{
			Username: cmd.Username,
			Email:    cmd.Email,
		}
		r.events = append(r.events, userRegisteredEvent)
	}
	return r
}

func (r RegisterCommandHandler) CollectEvents() []eventbus.Event {
	return r.events
}
