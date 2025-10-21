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
