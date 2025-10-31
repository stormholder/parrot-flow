package handlers

import (
	"context"
)

// CommandHandlerGroup groups related command handlers by domain concern
// This reduces constructor parameter explosion and improves cohesion
type CommandHandlerGroup[TCommand any, TResult any] interface {
	Handle(ctx context.Context, cmd TCommand) (TResult, error)
}

// QueryHandlerGroup groups related query handlers by domain concern
type QueryHandlerGroup[TQuery any, TResult any] interface {
	Handle(ctx context.Context, query TQuery) (TResult, error)
}

// HandlerRegistry provides a centralized way to store and retrieve handlers
// This eliminates the need for handlers to store 10+ dependencies as fields
type HandlerRegistry struct {
	commands map[string]interface{}
	queries  map[string]interface{}
}

func NewHandlerRegistry() *HandlerRegistry {
	return &HandlerRegistry{
		commands: make(map[string]interface{}),
		queries:  make(map[string]interface{}),
	}
}

func (r *HandlerRegistry) RegisterCommand(name string, handler interface{}) {
	r.commands[name] = handler
}

func (r *HandlerRegistry) RegisterQuery(name string, handler interface{}) {
	r.queries[name] = handler
}

func (r *HandlerRegistry) GetCommand(name string) interface{} {
	return r.commands[name]
}

func (r *HandlerRegistry) GetQuery(name string) interface{} {
	return r.queries[name]
}
