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
	var agents []*agent.Agent
	var err error

	// Filter by status
	if query.Status != nil {
		status, err := agent.NewAgentStatus(*query.Status)
		if err != nil {
			return nil, err
		}
		agents, err = h.repository.FindByStatus(ctx, status)
		if err != nil {
			return nil, err
		}
	} else if len(query.TagIDs) > 0 {
		// Filter by tags
		agents, err = h.repository.FindByTags(ctx, query.TagIDs)
		if err != nil {
			return nil, err
		}
	} else if query.BrowserType != nil {
		// Filter by browser type
		browserType, err := agent.NewBrowserType(*query.BrowserType)
		if err != nil {
			return nil, err
		}
		agents, err = h.repository.FindByBrowserType(ctx, browserType)
		if err != nil {
			return nil, err
		}
	} else if query.Platform != nil {
		// Filter by platform
		platform, err := agent.NewPlatform(*query.Platform)
		if err != nil {
			return nil, err
		}
		agents, err = h.repository.FindByPlatform(ctx, platform)
		if err != nil {
			return nil, err
		}
	} else {
		// No filters - return all agents
		agents, err = h.repository.FindByStatus(ctx, agent.AgentStatusOnline)
		if err != nil {
			return nil, err
		}
		// Get other statuses and merge
		idleAgents, _ := h.repository.FindByStatus(ctx, agent.AgentStatusIdle)
		busyAgents, _ := h.repository.FindByStatus(ctx, agent.AgentStatusBusy)
		disconnectedAgents, _ := h.repository.FindByStatus(ctx, agent.AgentStatusDisconnected)
		offlineAgents, _ := h.repository.FindByStatus(ctx, agent.AgentStatusOffline)

		agents = append(agents, idleAgents...)
		agents = append(agents, busyAgents...)
		agents = append(agents, disconnectedAgents...)
		agents = append(agents, offlineAgents...)
	}

	// Filter by health if requested
	if query.OnlyHealthy && query.HeartbeatTimeout > 0 {
		healthyAgents := make([]*agent.Agent, 0)
		for _, a := range agents {
			if a.IsHealthy(query.HeartbeatTimeout) {
				healthyAgents = append(healthyAgents, a)
			}
		}
		agents = healthyAgents
	}

	return agents, nil
}
