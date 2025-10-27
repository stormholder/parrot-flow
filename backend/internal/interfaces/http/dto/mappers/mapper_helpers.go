package mappers

import "time"

// FormatTimestamp converts a time.Time to RFC3339 string
func FormatTimestamp(t time.Time) string {
	return t.Format("2006-01-02T15:04:05Z07:00")
}

// FormatOptionalTimestamp converts optional time to string
func FormatOptionalTimestamp(t *time.Time) string {
	if t == nil {
		return ""
	}
	return t.Format("2006-01-02T15:04:05Z07:00")
}

// MapSlice applies a mapper function to a slice
func MapSlice[TIn any, TOut any](items []TIn, mapper func(TIn) TOut) []TOut {
	result := make([]TOut, len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}
	return result
}

// MapSlicePtr applies a mapper function to a slice of pointers
func MapSlicePtr[TIn any, TOut any](items []*TIn, mapper func(*TIn) TOut) []TOut {
	result := make([]TOut, len(items))
	for i, item := range items {
		result[i] = mapper(item)
	}
	return result
}

// Generic Response Builders - Eliminate mapper struct boilerplate

// CreateMapperFunc wraps a function to satisfy the Mapper interface
type CreateMapperFunc[TDomain any, TResponse any] func(TDomain) TResponse

func (f CreateMapperFunc[TDomain, TResponse]) Map(domain TDomain) TResponse {
	return f(domain)
}

// UpdateMapperFunc wraps a function to satisfy the Mapper interface
type UpdateMapperFunc[TDomain any, TResponse any] func(TDomain) TResponse

func (f UpdateMapperFunc[TDomain, TResponse]) Map(domain TDomain) TResponse {
	return f(domain)
}

// GetMapperFunc wraps a function to satisfy the Mapper interface
type GetMapperFunc[TDomain any, TResponse any] func(TDomain) TResponse

func (f GetMapperFunc[TDomain, TResponse]) Map(domain TDomain) TResponse {
	return f(domain)
}

// ListMapperFunc wraps a function to satisfy the Mapper interface for list operations
type ListMapperFunc[TDomain any, TResponse any] func([]*TDomain) TResponse

func (f ListMapperFunc[TDomain, TResponse]) Map(domains []*TDomain) TResponse {
	return f(domains)
}

// DeleteMapperFunc wraps a zero-argument function to satisfy mapper pattern
type DeleteMapperFunc[TResponse any] func() TResponse

func (f DeleteMapperFunc[TResponse]) Map() TResponse {
	return f()
}
