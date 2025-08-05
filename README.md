# GoCQRS

A Go implementation of the CQRS (Command Query Responsibility Segregation) pattern with CommandBus, QueryBus, and EventBus components.

## Overview

This library provides three core message buses:

- **CommandBus**: Handles write operations that modify system state
- **QueryBus**: Handles read operations that retrieve data
- **EventBus**: Handles domain event dispatch for decoupled communication

## Installation

```bash
go get github.com/avanboxel/gocqrs
```

## CommandBus

The CommandBus handles commands that modify system state and may produce domain events.

### Usage

```go
package main

import (
    "github.com/avanboxel/gocqrs"
)

// Define your command
type CreateUserCommand struct {
    Name  string
    Email string
}

// Implement CommandHandler
type CreateUserCommandHandler struct {
    events []gocqrs.Event
}

func (h *CreateUserCommandHandler) Handle(c gocqrs.Command) gocqrs.CommandHandler {
    cmd := c.(CreateUserCommand)
    
    // Business logic here
    // ...
    
    // Collect domain events
    h.events = append(h.events, UserCreatedEvent{Name: cmd.Name})
    
    return h
}

func (h *CreateUserCommandHandler) CollectEvents() []gocqrs.Event {
    return h.events
}

func main() {
    eventBus := gocqrs.DefaultSyncEventBus() // or DefaultAsyncEventBus() for concurrent execution
    commandBus := gocqrs.DefaultCommandBus(eventBus)
    
    // Register handler
    commandBus.Register(CreateUserCommand{}, &CreateUserCommandHandler{})
    
    // Execute command synchronously
    commandBus.Execute(CreateUserCommand{Name: "John", Email: "john@example.com"})
    
    // Or dispatch asynchronously
    commandBus.Dispatch(CreateUserCommand{Name: "Jane", Email: "jane@example.com"})
}
```

## QueryBus

The QueryBus handles read operations that retrieve data without modifying system state.

### Usage

```go
package main

import (
    "github.com/avanboxel/gocqrs"
)

// Define your query
type GetUserQuery struct {
    ID int
}

// Implement QueryHandler
type GetUserQueryHandler struct{}

func (h *GetUserQueryHandler) Handle(q gocqrs.Query) gocqrs.QueryResult {
    query := q.(GetUserQuery)
    
    // Fetch data logic here
    user := fetchUser(query.ID)
    
    if user != nil {
        return gocqrs.QueryResult{
            Payload: user,
            Success: true,
        }
    }
    
    return gocqrs.QueryResult{
        Payload: "User not found",
        Success: false,
    }
}

func main() {
    queryBus := gocqrs.DefaultQueryBus()
    
    // Register handler
    queryBus.Register(GetUserQuery{}, &GetUserQueryHandler{})
    
    // Execute query
    result := queryBus.Ask(GetUserQuery{ID: 1})
    
    if result.Success {
        user := result.Payload.(User)
        // Handle successful result
    } else {
        // Handle failure
        errorMsg := result.Payload.(string)
    }
}
```

## EventBus

The EventBus handles domain event dispatch for decoupled communication between system components.

### Usage

```go
package main

import (
    "github.com/avanboxel/gocqrs"
)

// Define your event
type UserCreatedEvent struct {
    Name string
}

func (e UserCreatedEvent) GetEventType() string {
    return "UserCreated"
}

// Define event handler
func userCreatedHandler(e gocqrs.Event) {
    event := e.(UserCreatedEvent)
    // Handle the event (send email, update cache, etc.)
    sendWelcomeEmail(event.Name)
}

func main() {
    eventBus := gocqrs.DefaultSyncEventBus() // or DefaultAsyncEventBus() for concurrent execution
    
    // Register event handler
    eventBus.Register("UserCreated", userCreatedHandler)
    
    // Dispatch event
    eventBus.Dispatch(UserCreatedEvent{Name: "John"})
}
```

## Complete Example

See the [examples](./examples/) directory for complete working examples:

- [get-username](./examples/get-username/) - QueryBus example
- [user-register](./examples/user-register/) - CommandBus example
- [validate-username](./examples/validate-username/) - Additional example

## Key Features

- **Type Safety**: Uses Go's type system with reflection for handler registration
- **CQRS Pattern**: Clear separation between commands (write) and queries (read)
- **Event Sourcing**: Commands can produce domain events
- **Synchronous & Asynchronous**: CommandBus supports both execution modes
- **Error Handling**: QueryBus returns structured results with success indicators
- **Decoupled Architecture**: EventBus enables loose coupling between components

## Best Practices

1. **Commands**: Should be imperative (CreateUser, UpdateProduct)
2. **Queries**: Should be noun-based (GetUser, FindProducts)
3. **Events**: Should be past tense (UserCreated, ProductUpdated)
4. **Handlers**: Keep them focused and single-purpose
5. **Error Handling**: Use QueryResult.Success for query error handling

## License

MIT License - see [LICENSE](LICENSE) file for details.