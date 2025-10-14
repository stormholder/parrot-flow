package agent

import (
	"context"
	"time"

	"parrotflow/internal/domain/agent"
)

type GetStaleAgentsQuery struct {
	HeartbeatTimeout time.Duration
}

type GetStaleAgentsQueryHandler struct {
	repository agent.Repository
}

func NewGetStaleAgentsQueryHandler(repository agent.Repository) *GetStaleAgentsQueryHandler {
	return &GetStaleAgentsQueryHandler{
		repository: repository,
	}
}

func (h *GetStaleAgentsQueryHandler) Handle(ctx context.Context, query GetStaleAgentsQuery) ([]*agent.Agent, error) {
	return h.repository.FindStaleAgents(ctx, query.HeartbeatTimeout)
}
