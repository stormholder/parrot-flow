package command

import (
	"context"
	"errors"
	"testing"

	"parrotflow/internal/domain/tag"
	"parrotflow/internal/domain/shared"
)

// MockTagRepository is a mock implementation of tag.Repository for testing
type MockTagRepository struct {
	SaveFunc       func(ctx context.Context, t *tag.Tag) error
	ExistsFunc     func(ctx context.Context, name string) (bool, error)
	FindByIDFunc   func(ctx context.Context, id tag.TagID) (*tag.Tag, error)
}

func (m *MockTagRepository) Save(ctx context.Context, t *tag.Tag) error {
	if m.SaveFunc != nil {
		return m.SaveFunc(ctx, t)
	}
	return nil
}

func (m *MockTagRepository) Exists(ctx context.Context, name string) (bool, error) {
	if m.ExistsFunc != nil {
		return m.ExistsFunc(ctx, name)
	}
	return false, nil
}

func (m *MockTagRepository) FindByID(ctx context.Context, id tag.TagID) (*tag.Tag, error) {
	if m.FindByIDFunc != nil {
		return m.FindByIDFunc(ctx, id)
	}
	return nil, errors.New("not implemented")
}

func (m *MockTagRepository) FindByName(ctx context.Context, name string) (*tag.Tag, error) {
	return nil, errors.New("not implemented")
}

func (m *MockTagRepository) FindByCategory(ctx context.Context, category tag.TagCategory) ([]*tag.Tag, error) {
	return nil, errors.New("not implemented")
}

func (m *MockTagRepository) FindAll(ctx context.Context) ([]*tag.Tag, error) {
	return nil, errors.New("not implemented")
}

func (m *MockTagRepository) FindByIDs(ctx context.Context, ids []tag.TagID) ([]*tag.Tag, error) {
	return nil, errors.New("not implemented")
}

func (m *MockTagRepository) Delete(ctx context.Context, id tag.TagID) error {
	return errors.New("not implemented")
}

// MockEventBus is a mock implementation of shared.EventBus for testing
type MockEventBus struct {
	PublishedEvents []shared.DomainEvent
}

func (m *MockEventBus) Publish(event shared.DomainEvent) error {
	m.PublishedEvents = append(m.PublishedEvents, event)
	return nil
}

func (m *MockEventBus) Subscribe(handler shared.EventHandler) error {
	return nil
}

// TestCreateTagCommand_Success demonstrates DI benefits - easy mocking
func TestCreateTagCommand_Success(t *testing.T) {
	// Arrange - inject mocks
	mockRepo := &MockTagRepository{
		ExistsFunc: func(ctx context.Context, name string) (bool, error) {
			return false, nil // Tag doesn't exist
		},
		SaveFunc: func(ctx context.Context, tag *tag.Tag) error {
			return nil // Save succeeds
		},
	}
	mockBus := &MockEventBus{}

	handler := NewCreateTagCommandHandler(mockRepo, mockBus)

	cmd := CreateTagCommand{
		Name:        "test-tag",
		Category:    "custom",
		Description: "Test tag",
		Color:       "#FF0000",
	}

	// Act
	result, err := handler.Handle(context.Background(), cmd)

	// Assert
	if err != nil {
		t.Fatalf("Expected no error, got: %v", err)
	}

	if result.Name != "test-tag" {
		t.Errorf("Expected tag name 'test-tag', got: %s", result.Name)
	}

	if len(mockBus.PublishedEvents) != 1 {
		t.Errorf("Expected 1 event published, got: %d", len(mockBus.PublishedEvents))
	}
}

// TestCreateTagCommand_AlreadyExists demonstrates handling duplicate tags
func TestCreateTagCommand_AlreadyExists(t *testing.T) {
	// Arrange - tag already exists
	mockRepo := &MockTagRepository{
		ExistsFunc: func(ctx context.Context, name string) (bool, error) {
			return true, nil // Tag exists
		},
	}
	mockBus := &MockEventBus{}

	handler := NewCreateTagCommandHandler(mockRepo, mockBus)

	cmd := CreateTagCommand{
		Name:     "existing-tag",
		Category: "custom",
	}

	// Act
	result, err := handler.Handle(context.Background(), cmd)

	// Assert
	if err != tag.ErrTagAlreadyExists {
		t.Errorf("Expected ErrTagAlreadyExists, got: %v", err)
	}

	if result != nil {
		t.Error("Expected nil result when tag exists")
	}

	if len(mockBus.PublishedEvents) != 0 {
		t.Error("Expected no events when tag already exists")
	}
}

// TestCreateTagCommand_RepositoryError demonstrates error handling
func TestCreateTagCommand_RepositoryError(t *testing.T) {
	// Arrange - repository fails
	expectedErr := errors.New("database connection lost")
	mockRepo := &MockTagRepository{
		ExistsFunc: func(ctx context.Context, name string) (bool, error) {
			return false, expectedErr
		},
	}
	mockBus := &MockEventBus{}

	handler := NewCreateTagCommandHandler(mockRepo, mockBus)

	cmd := CreateTagCommand{
		Name:     "test-tag",
		Category: "custom",
	}

	// Act
	result, err := handler.Handle(context.Background(), cmd)

	// Assert
	if err != expectedErr {
		t.Errorf("Expected repository error, got: %v", err)
	}

	if result != nil {
		t.Error("Expected nil result on repository error")
	}
}
