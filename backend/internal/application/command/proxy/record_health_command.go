package proxy

import (
	"context"
	"time"

	"parrotflow/internal/domain/proxy"
	"parrotflow/pkg/shared"
)

type RecordHealthCommand struct {
	ID         string
	Success    bool
	LatencyMs  int
	ErrorMsg   string
}

type RecordHealthCommandHandler struct {
	repository proxy.ProxyRepository
	eventBus   shared.EventBus
}

func NewRecordHealthCommandHandler(repository proxy.ProxyRepository, eventBus shared.EventBus) *RecordHealthCommandHandler {
	return &RecordHealthCommandHandler{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (h *RecordHealthCommandHandler) Handle(ctx context.Context, cmd RecordHealthCommand) (*proxy.Proxy, error) {
	// Parse proxy ID
	proxyID, err := proxy.NewProxyID(cmd.ID)
	if err != nil {
		return nil, err
	}

	// Find proxy
	p, err := h.repository.FindByID(ctx, proxyID)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, proxy.ErrProxyNotFound
	}

	// Record success or failure
	if cmd.Success {
		p.RecordSuccess(time.Duration(cmd.LatencyMs) * time.Millisecond)
	} else {
		p.RecordFailure(cmd.ErrorMsg)
	}

	// Save to repository
	if err := h.repository.Save(ctx, p); err != nil {
		return nil, err
	}

	// Publish domain events
	for _, event := range p.Events {
		h.eventBus.Publish(event)
	}

	return p, nil
}
