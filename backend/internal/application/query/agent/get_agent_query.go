package agent

import (
	"context"

	"parrotflow/internal/domain/agent"
)

type GetAgentQuery struct {
	ID agent.AgentID
}

type GetAgentQueryHandler struct {
	repository agent.Repository
}

func NewGetAgentQueryHandler(repository agent.Repository) *GetAgentQueryHandler {
	return &GetAgentQueryHandler{
		repository: repository,
	}
}

func (h *GetAgentQueryHandler) Handle(ctx context.Context, query GetAgentQuery) (*agent.Agent, error) {
	return h.repository.FindByID(ctx, query.ID)
}
