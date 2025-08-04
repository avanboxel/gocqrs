package validateusername

import (
	"fmt"

	"github.com/avanboxel/gocqrs"
)

type ValidateUsernameCommand struct {
	Username string
}

type UsernameValidated struct {
	Username string
	IsValid  bool
	Reason   string
}

func (e UsernameValidated) GetEventType() string {
	return "UsernameValidated"
}

type ValidateUsernameHandler struct {
	events []gocqrs.Event
}

func (v *ValidateUsernameHandler) Handle(c gocqrs.Command) gocqrs.CommandHandler {
	if cmd, ok := c.(ValidateUsernameCommand); ok {
		fmt.Printf("Validating username: %s\n", cmd.Username)

		var isValid bool
		var reason string

		usernameLen := len(cmd.Username)
		if usernameLen >= 8 && usernameLen <= 16 {
			isValid = true
			reason = "Username is valid"
			fmt.Println("Username validation passed")
		} else {
			isValid = false
			if usernameLen < 8 {
				reason = "Username too short (minimum 8 characters)"
			} else {
				reason = "Username too long (maximum 16 characters)"
			}
			fmt.Printf("Username validation failed: %s\n", reason)
		}

		validationEvent := UsernameValidated{
			Username: cmd.Username,
			IsValid:  isValid,
			Reason:   reason,
		}
		v.events = append(v.events, validationEvent)
	}
	return v
}

func (v *ValidateUsernameHandler) CollectEvents() []gocqrs.Event {
	return v.events
}
