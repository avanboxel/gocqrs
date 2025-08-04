// Package querybus provides a query bus implementation for CQRS (Command Query Responsibility Segregation) pattern.
// The query bus handles read operations that return data without modifying system state.
package querybus

// Query represents any query object that can be handled by a QueryHandler.
// Queries are read-only operations that retrieve data from the system.
// Examples: GetUserQuery, FindProductsQuery, CountOrdersQuery
type Query any

// QueryResult represents the result of a query operation.
// It contains the actual data payload and a success indicator.
type QueryResult struct {
	// Payload contains the actual query result data.
	// The type depends on the specific query being executed.
	Payload any
	
	// Success indicates whether the query was executed successfully.
	// When false, Payload may contain error information.
	Success bool
}

// QueryHandler defines the interface for handling queries.
// Each query type should have a corresponding handler that implements this interface.
type QueryHandler interface {
	// Handle processes the given query and returns a QueryResult.
	// The result contains both the payload and success status.
	// Handlers should set Success to false if the query fails.
	Handle(q Query) QueryResult
}