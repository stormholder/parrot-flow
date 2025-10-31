package mappers

import (
	"testing"
	"time"
)

func TestFormatTimestamp(t *testing.T) {
	// Test case: Known time
	testTime := time.Date(2024, 10, 27, 15, 30, 45, 0, time.UTC)
	result := FormatTimestamp(testTime)

	expected := "2024-10-27T15:30:45Z"
	if result != expected {
		t.Errorf("FormatTimestamp() = %v, want %v", result, expected)
	}
}

func TestMapSlice(t *testing.T) {
	// Test case: Map integers to strings
	input := []int{1, 2, 3, 4, 5}
	mapper := func(i int) string {
		return string(rune('0' + i))
	}

	result := MapSlice(input, mapper)

	if len(result) != len(input) {
		t.Errorf("MapSlice() length = %v, want %v", len(result), len(input))
	}

	expected := []string{"1", "2", "3", "4", "5"}
	for i, v := range result {
		if v != expected[i] {
			t.Errorf("MapSlice()[%d] = %v, want %v", i, v, expected[i])
		}
	}
}

func TestMapSlicePtr(t *testing.T) {
	// Test case: Map pointers
	type TestStruct struct {
		Value int
	}

	input := []*TestStruct{{1}, {2}, {3}}
	mapper := func(ts *TestStruct) int {
		return ts.Value * 2
	}

	result := MapSlicePtr(input, mapper)

	expected := []int{2, 4, 6}
	if len(result) != len(expected) {
		t.Errorf("MapSlicePtr() length = %v, want %v", len(result), len(expected))
	}

	for i, v := range result {
		if v != expected[i] {
			t.Errorf("MapSlicePtr()[%d] = %v, want %v", i, v, expected[i])
		}
	}
}

func TestMapSlicePtr_EmptySlice(t *testing.T) {
	// Test case: Empty input
	type TestStruct struct {
		Value int
	}

	input := []*TestStruct{}
	mapper := func(ts *TestStruct) int {
		return ts.Value
	}

	result := MapSlicePtr(input, mapper)

	if len(result) != 0 {
		t.Errorf("MapSlicePtr() with empty input should return empty slice, got length %v", len(result))
	}
}

func TestCreateMapperFunc(t *testing.T) {
	// Test case: Mapper function wrapping
	type Domain struct {
		Name string
	}
	type Response struct {
		Output string
	}

	mapperFunc := func(d *Domain) *Response {
		return &Response{Output: "Hello, " + d.Name}
	}

	mapper := CreateMapperFunc[*Domain, *Response](mapperFunc)

	input := &Domain{Name: "World"}
	result := mapper.Map(input)

	if result.Output != "Hello, World" {
		t.Errorf("CreateMapperFunc().Map() = %v, want %v", result.Output, "Hello, World")
	}
}
