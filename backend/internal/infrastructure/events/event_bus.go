package events

import (
	"log"
	"parrotflow/internal/domain/shared"
	"sync"
)

type InMemoryEventBus struct {
	handlers []shared.EventHandler
	mu       sync.RWMutex
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		handlers: make([]shared.EventHandler, 0),
	}
}

func (bus *InMemoryEventBus) Publish(event shared.DomainEvent) error {
	bus.mu.RLock()
	defer bus.mu.RUnlock()

	for _, handler := range bus.handlers {
		if handler.CanHandle(event.EventType()) {
			if err := handler.Handle(event); err != nil {
				log.Printf("Error handling event %s: %v", event.EventType(), err)
			}
		}
	}

	return nil
}

func (bus *InMemoryEventBus) Subscribe(handler shared.EventHandler) error {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	bus.handlers = append(bus.handlers, handler)
	return nil
}

type AsyncEventBus struct {
	handlers []shared.EventHandler
	mu       sync.RWMutex
}

func NewAsyncEventBus() *AsyncEventBus {
	return &AsyncEventBus{
		handlers: make([]shared.EventHandler, 0),
	}
}

func (bus *AsyncEventBus) Publish(event shared.DomainEvent) error {
	bus.mu.RLock()
	defer bus.mu.RUnlock()

	for _, handler := range bus.handlers {
		if handler.CanHandle(event.EventType()) {
			go func(h shared.EventHandler, e shared.DomainEvent) {
				if err := h.Handle(e); err != nil {
					log.Printf("Error handling event %s: %v", e.EventType(), err)
				}
			}(handler, event)
		}
	}

	return nil
}

func (bus *AsyncEventBus) Subscribe(handler shared.EventHandler) error {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	bus.handlers = append(bus.handlers, handler)
	return nil
}
