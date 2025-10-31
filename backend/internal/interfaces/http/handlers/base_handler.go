package handlers

import (
	"context"
	"fmt"
)

type CommandHandler[TCommand any, TResult any] interface {
	Handle(ctx context.Context, cmd TCommand) (TResult, error)
}

type QueryHandler[TQuery any, TResult any] interface {
	Handle(ctx context.Context, query TQuery) (TResult, error)
}

type Mapper[TDomain any, TDTO any] interface {
	Map(domain TDomain) TDTO
}

type IDParser[TID any] func(idStr string) (TID, error)

func HandleCommand[TRequest any, TCommand any, TDomain any, TResponse any](
	ctx context.Context,
	request TRequest,
	buildCommand func(TRequest) (TCommand, error),
	handler CommandHandler[TCommand, TDomain],
	mapper Mapper[TDomain, TResponse],
) (TResponse, error) {
	var zero TResponse

	cmd, err := buildCommand(request)
	if err != nil {
		return zero, fmt.Errorf("invalid request: %w", err)
	}

	result, err := handler.Handle(ctx, cmd)
	if err != nil {
		return zero, err
	}

	return mapper.Map(result), nil
}

func HandleQuery[TRequest any, TQuery any, TDomain any, TResponse any](
	ctx context.Context,
	request TRequest,
	buildQuery func(TRequest) (TQuery, error),
	handler QueryHandler[TQuery, TDomain],
	mapper Mapper[TDomain, TResponse],
) (TResponse, error) {
	var zero TResponse

	query, err := buildQuery(request)
	if err != nil {
		return zero, fmt.Errorf("invalid request: %w", err)
	}

	result, err := handler.Handle(ctx, query)
	if err != nil {
		return zero, err
	}

	return mapper.Map(result), nil
}

type SimpleCommandHandler[TCommand any] interface {
	Handle(ctx context.Context, cmd TCommand) error
}

type SimpleCommandHandlerFunc[TCommand any] func(context.Context, TCommand) error

func (f SimpleCommandHandlerFunc[TCommand]) Handle(ctx context.Context, cmd TCommand) error {
	return f(ctx, cmd)
}

func HandleSimpleCommand[TRequest any, TCommand any, TResponse any](
	ctx context.Context,
	request TRequest,
	buildCommand func(TRequest) (TCommand, error),
	handler SimpleCommandHandler[TCommand],
	buildResponse func() TResponse,
) (TResponse, error) {
	var zero TResponse

	cmd, err := buildCommand(request)
	if err != nil {
		return zero, fmt.Errorf("invalid request: %w", err)
	}

	err = handler.Handle(ctx, cmd)
	if err != nil {
		return zero, err
	}

	return buildResponse(), nil
}

type CommandHandlerFunc[TCommand any, TResult any] func(context.Context, TCommand) (TResult, error)

func (f CommandHandlerFunc[TCommand, TResult]) Handle(ctx context.Context, cmd TCommand) (TResult, error) {
	return f(ctx, cmd)
}

type QueryHandlerFunc[TQuery any, TResult any] func(context.Context, TQuery) (TResult, error)

func (f QueryHandlerFunc[TQuery, TResult]) Handle(ctx context.Context, query TQuery) (TResult, error) {
	return f(ctx, query)
}

type MapperFunc[TDomain any, TDTO any] func(TDomain) TDTO

func (f MapperFunc[TDomain, TDTO]) Map(domain TDomain) TDTO {
	return f(domain)
}
