package shared

import "strconv"

func ParseInt(value string, defaultValue int) int {
	if value == "" {
		return defaultValue
	}
	if result, err := strconv.Atoi(value); err == nil {
		return result
	}
	return defaultValue
}

func ParseString(value string, defaultValue string) string {
	if value == "" {
		return defaultValue
	}
	return value
}
