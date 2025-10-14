package agent

import (
	"context"

	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/tag"
)

type GetAvailableAgentsQuery struct {
	TagIDs      []tag.TagID
	BrowserType *string
	Platform    *string
}

type GetAvailableAgentsQueryHandler struct {
	repository agent.Repository
}

func NewGetAvailableAgentsQueryHandler(repository agent.Repository) *GetAvailableAgentsQueryHandler {
	return &GetAvailableAgentsQueryHandler{
		repository: repository,
	}
}

func (h *GetAvailableAgentsQueryHandler) Handle(ctx context.Context, query GetAvailableAgentsQuery) ([]*agent.Agent, error) {
	// Start with all available agents
	agents, err := h.repository.FindAvailable(ctx)
	if err != nil {
		return nil, err
	}

	// Filter by tags if provided
	if len(query.TagIDs) > 0 {
		filteredAgents := make([]*agent.Agent, 0)
		for _, a := range agents {
			if a.HasAllTags(query.TagIDs) {
				filteredAgents = append(filteredAgents, a)
			}
		}
		agents = filteredAgents
	}

	// Filter by browser type if provided
	if query.BrowserType != nil {
		browserType, err := agent.NewBrowserType(*query.BrowserType)
		if err != nil {
			return nil, err
		}
		filteredAgents := make([]*agent.Agent, 0)
		for _, a := range agents {
			if a.Capabilities.HasBrowser(browserType) {
				filteredAgents = append(filteredAgents, a)
			}
		}
		agents = filteredAgents
	}

	// Filter by platform if provided
	if query.Platform != nil {
		platform, err := agent.NewPlatform(*query.Platform)
		if err != nil {
			return nil, err
		}
		filteredAgents := make([]*agent.Agent, 0)
		for _, a := range agents {
			if a.Capabilities.OS.Platform.String() == platform.String() {
				filteredAgents = append(filteredAgents, a)
			}
		}
		agents = filteredAgents
	}

	return agents, nil
}
