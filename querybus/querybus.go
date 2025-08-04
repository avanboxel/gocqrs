package querybus

import "reflect"

// QueryBus defines the interface for a query bus that handles read operations.
// It provides methods to execute queries and register query handlers.
type QueryBus interface {
	// Ask executes a query synchronously and returns the result immediately.
	// It looks up the registered handler for the query type and delegates execution.
	// Panics if no handler is registered for the query type.
	Ask(q Query) QueryResult
	
	// Register associates a query type with its corresponding handler.
	// The query parameter is used to determine the type name for registration.
	// Only one handler can be registered per query type (last registration wins).
	Register(q Query, qh QueryHandler)
}

// defaultQueryBus is the default implementation of QueryBus.
// It uses reflection to map query types to their handlers.
type defaultQueryBus struct {
	// handlers maps query type names to their corresponding handlers
	handlers map[string]QueryHandler
}

// Ask executes the given query by finding its registered handler.
// It uses reflection to determine the query type name and looks up the handler.
// Panics if no handler is registered for the query type.
func (d *defaultQueryBus) Ask(q Query) QueryResult {
	qh := d.handlers[reflect.TypeOf(q).Name()]
	if qh == nil {
		panic("no handler registered for query type: " + reflect.TypeOf(q).Name())
	}
	return qh.Handle(q)
}

// Register stores a query handler for the given query type.
// It uses reflection to extract the type name from the query instance.
// If a handler already exists for this query type, it will be replaced.
func (d *defaultQueryBus) Register(q Query, qh QueryHandler) {
	d.handlers[reflect.TypeOf(q).Name()] = qh
}

// NewDefault creates a new instance of the default query bus implementation.
// Returns a QueryBus that uses reflection-based handler lookup.
func NewDefault() *defaultQueryBus {
	return &defaultQueryBus{
		handlers: make(map[string]QueryHandler),
	}
}