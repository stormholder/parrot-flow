package agent

import (
	"context"
	"time"

	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/tag"
)

type ListAgentsQuery struct {
	Status       *string
	TagIDs       []tag.TagID
	BrowserType  *string
	Platform     *string
	OnlyHealthy  bool
	HeartbeatTimeout time.Duration
}

type ListAgentsQueryHandler struct {
	repository agent.Repository
}

func NewListAgentsQueryHandler(repository agent.Repository) *ListAgentsQueryHandler {
	return &ListAgentsQueryHandler{
		repository: repository,
	}
}

func (h *ListAgentsQueryHandler) Handle(ctx context.Context, query ListAgentsQuery) ([]*agent.Agent, error) {
	// Build search criteria from query
	criteria := agent.SearchCriteria{
		TagIDs:           query.TagIDs,
		OnlyHealthy:      query.OnlyHealthy,
		HeartbeatTimeout: query.HeartbeatTimeout,
	}

	// Parse and validate status if provided
	if query.Status != nil {
		status, err := agent.NewAgentStatus(*query.Status)
		if err != nil {
			return nil, err
		}
		criteria.Status = &status
	}

	// Parse and validate browser type if provided
	if query.BrowserType != nil {
		browserType, err := agent.NewBrowserType(*query.BrowserType)
		if err != nil {
			return nil, err
		}
		criteria.BrowserType = &browserType
	}

	// Parse and validate platform if provided
	if query.Platform != nil {
		platform, err := agent.NewPlatform(*query.Platform)
		if err != nil {
			return nil, err
		}
		criteria.Platform = &platform
	}

	// Single repository call handles all filtering with combined criteria
	return h.repository.FindByCriteria(ctx, criteria)
}
