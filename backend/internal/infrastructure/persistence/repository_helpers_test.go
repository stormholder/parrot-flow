package persistence

import (
	"errors"
	"testing"
)

// Test models and entities
type TestModel struct {
	ID   int
	Name string
}

type TestEntity struct {
	ID   int
	Name string
}

func TestConvertSliceToDomainPtr_Success(t *testing.T) {
	// Test case: Successful conversion
	models := []TestModel{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
		{ID: 3, Name: "Charlie"},
	}

	converter := func(m *TestModel) (*TestEntity, error) {
		return &TestEntity{ID: m.ID, Name: m.Name}, nil
	}

	result, err := ConvertSliceToDomainPtr(models, converter)

	if err != nil {
		t.Errorf("ConvertSliceToDomainPtr() error = %v, want nil", err)
	}

	if len(result) != 3 {
		t.Errorf("ConvertSliceToDomainPtr() length = %v, want 3", len(result))
	}

	// Verify first entity
	if result[0].Name != "Alice" {
		t.Errorf("ConvertSliceToDomainPtr()[0].Name = %v, want Alice", result[0].Name)
	}
}

func TestConvertSliceToDomainPtr_EmptySlice(t *testing.T) {
	// Test case: Empty input slice
	models := []TestModel{}

	converter := func(m *TestModel) (*TestEntity, error) {
		return &TestEntity{ID: m.ID, Name: m.Name}, nil
	}

	result, err := ConvertSliceToDomainPtr(models, converter)

	if err != nil {
		t.Errorf("ConvertSliceToDomainPtr() error = %v, want nil", err)
	}

	if len(result) != 0 {
		t.Errorf("ConvertSliceToDomainPtr() with empty input should return empty slice, got length %v", len(result))
	}
}

func TestConvertSliceToDomainPtr_ConversionError(t *testing.T) {
	// Test case: Converter returns error
	models := []TestModel{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	converter := func(m *TestModel) (*TestEntity, error) {
		if m.ID == 2 {
			return nil, errors.New("conversion failed for ID 2")
		}
		return &TestEntity{ID: m.ID, Name: m.Name}, nil
	}

	result, err := ConvertSliceToDomainPtr(models, converter)

	if err == nil {
		t.Error("ConvertSliceToDomainPtr() should return error when converter fails")
	}

	if result != nil {
		t.Errorf("ConvertSliceToDomainPtr() should return nil result on error, got %v", result)
	}
}

func TestConvertSliceToDomain_NonPointer(t *testing.T) {
	// Test case: Non-pointer conversion
	models := []TestModel{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	converter := func(m *TestModel) (TestEntity, error) {
		return TestEntity{ID: m.ID, Name: m.Name}, nil
	}

	result, err := ConvertSliceToDomain(models, converter)

	if err != nil {
		t.Errorf("ConvertSliceToDomain() error = %v, want nil", err)
	}

	if len(result) != 2 {
		t.Errorf("ConvertSliceToDomain() length = %v, want 2", len(result))
	}

	if result[0].Name != "Alice" {
		t.Errorf("ConvertSliceToDomain()[0].Name = %v, want Alice", result[0].Name)
	}
}
