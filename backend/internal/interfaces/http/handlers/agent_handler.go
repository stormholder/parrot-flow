package handlers

import (
	"context"
	"fmt"
	"strings"
	"time"

	agentcommand "parrotflow/internal/application/command/agent"
	agentquery "parrotflow/internal/application/query/agent"
	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/mappers"
	"parrotflow/internal/interfaces/http/dto/queries"
)

type AgentHandler struct {
	// Command handlers
	registerCommandHandler     *agentcommand.RegisterAgentCommandHandler
	updateHeartbeatCommandHandler *agentcommand.UpdateHeartbeatCommandHandler
	assignRunCommandHandler    *agentcommand.AssignRunCommandHandler
	releaseRunCommandHandler   *agentcommand.ReleaseRunCommandHandler
	updateCommandHandler       *agentcommand.UpdateAgentCommandHandler
	deregisterCommandHandler   *agentcommand.DeregisterAgentCommandHandler

	// Query handlers
	getQueryHandler            *agentquery.GetAgentQueryHandler
	listQueryHandler           *agentquery.ListAgentsQueryHandler
	getAvailableQueryHandler   *agentquery.GetAvailableAgentsQueryHandler
	getStaleQueryHandler       *agentquery.GetStaleAgentsQueryHandler
}

func NewAgentHandler(
	registerCommandHandler *agentcommand.RegisterAgentCommandHandler,
	updateHeartbeatCommandHandler *agentcommand.UpdateHeartbeatCommandHandler,
	assignRunCommandHandler *agentcommand.AssignRunCommandHandler,
	releaseRunCommandHandler *agentcommand.ReleaseRunCommandHandler,
	updateCommandHandler *agentcommand.UpdateAgentCommandHandler,
	deregisterCommandHandler *agentcommand.DeregisterAgentCommandHandler,
	getQueryHandler *agentquery.GetAgentQueryHandler,
	listQueryHandler *agentquery.ListAgentsQueryHandler,
	getAvailableQueryHandler *agentquery.GetAvailableAgentsQueryHandler,
	getStaleQueryHandler *agentquery.GetStaleAgentsQueryHandler,
) *AgentHandler {
	return &AgentHandler{
		registerCommandHandler:     registerCommandHandler,
		updateHeartbeatCommandHandler: updateHeartbeatCommandHandler,
		assignRunCommandHandler:    assignRunCommandHandler,
		releaseRunCommandHandler:   releaseRunCommandHandler,
		updateCommandHandler:       updateCommandHandler,
		deregisterCommandHandler:   deregisterCommandHandler,
		getQueryHandler:            getQueryHandler,
		listQueryHandler:           listQueryHandler,
		getAvailableQueryHandler:   getAvailableQueryHandler,
		getStaleQueryHandler:       getStaleQueryHandler,
	}
}

// RegisterAgent handles agent registration
func (h *AgentHandler) RegisterAgent(ctx context.Context, input *commands.RegisterAgentRequest) (*commands.RegisterAgentResponse, error) {
	// Convert DTOs to domain objects
	capabilities, err := mappers.FromCapabilitiesDTO(input.Body.Capabilities)
	if err != nil {
		return nil, fmt.Errorf("invalid capabilities: %w", err)
	}

	connectionInfo, err := mappers.FromConnectionInfoDTO(input.Body.ConnectionInfo)
	if err != nil {
		return nil, fmt.Errorf("invalid connection info: %w", err)
	}

	// Create command
	cmd := agentcommand.RegisterAgentCommand{
		Name:           input.Body.Name,
		Capabilities:   capabilities,
		ConnectionInfo: connectionInfo,
	}

	// Execute command
	a, err := h.registerCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	// Convert to response
	return mappers.ToRegisterAgentResponse(a), nil
}

// UpdateHeartbeat handles agent heartbeat updates
func (h *AgentHandler) UpdateHeartbeat(ctx context.Context, input *commands.UpdateHeartbeatRequest) (*commands.UpdateHeartbeatResponse, error) {
	// Parse agent ID
	agentID, err := agent.NewAgentID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid agent ID: %w", err)
	}

	// Create command
	cmd := agentcommand.UpdateHeartbeatCommand{
		AgentID: agentID,
	}

	// Execute command
	a, err := h.updateHeartbeatCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	// Convert to response
	return mappers.ToUpdateHeartbeatResponse(a), nil
}

// AssignRun handles assigning a run to an agent
func (h *AgentHandler) AssignRun(ctx context.Context, input *commands.AssignRunRequest) (*commands.AssignRunResponse, error) {
	// Parse agent ID
	agentID, err := agent.NewAgentID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid agent ID: %w", err)
	}

	// Create command
	cmd := agentcommand.AssignRunCommand{
		AgentID: agentID,
	}

	// Execute command
	a, err := h.assignRunCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	// Convert to response
	return mappers.ToAssignRunResponse(a), nil
}

// ReleaseRun handles releasing a run from an agent
func (h *AgentHandler) ReleaseRun(ctx context.Context, input *commands.ReleaseRunRequest) (*commands.ReleaseRunResponse, error) {
	// Parse agent ID
	agentID, err := agent.NewAgentID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid agent ID: %w", err)
	}

	// Create command
	cmd := agentcommand.ReleaseRunCommand{
		AgentID: agentID,
	}

	// Execute command
	a, err := h.releaseRunCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	// Convert to response
	return mappers.ToReleaseRunResponse(a), nil
}

// UpdateAgent handles agent updates
func (h *AgentHandler) UpdateAgent(ctx context.Context, input *commands.UpdateAgentRequest) (*commands.UpdateAgentResponse, error) {
	// Parse agent ID
	agentID, err := agent.NewAgentID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid agent ID: %w", err)
	}

	// Create command
	cmd := agentcommand.UpdateAgentCommand{
		AgentID: agentID,
		Name:    input.Body.Name,
	}

	// Convert capabilities if provided
	if input.Body.Capabilities != nil {
		capabilities, err := mappers.FromCapabilitiesDTO(*input.Body.Capabilities)
		if err != nil {
			return nil, fmt.Errorf("invalid capabilities: %w", err)
		}
		cmd.Capabilities = &capabilities
	}

	// Convert tag IDs to add
	if len(input.Body.TagsToAdd) > 0 {
		cmd.TagsToAdd = make([]tag.TagID, len(input.Body.TagsToAdd))
		for i, tagIDStr := range input.Body.TagsToAdd {
			tagID, err := tag.NewTagID(tagIDStr)
			if err != nil {
				return nil, fmt.Errorf("invalid tag ID: %w", err)
			}
			cmd.TagsToAdd[i] = tagID
		}
	}

	// Convert tag IDs to remove
	if len(input.Body.TagsToRemove) > 0 {
		cmd.TagsToRemove = make([]tag.TagID, len(input.Body.TagsToRemove))
		for i, tagIDStr := range input.Body.TagsToRemove {
			tagID, err := tag.NewTagID(tagIDStr)
			if err != nil {
				return nil, fmt.Errorf("invalid tag ID: %w", err)
			}
			cmd.TagsToRemove[i] = tagID
		}
	}

	// Execute command
	a, err := h.updateCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	// Convert to response
	return mappers.ToUpdateAgentResponse(a), nil
}

// DeregisterAgent handles agent deregistration
func (h *AgentHandler) DeregisterAgent(ctx context.Context, input *commands.DeregisterAgentRequest) (*commands.DeregisterAgentResponse, error) {
	// Parse agent ID
	agentID, err := agent.NewAgentID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid agent ID: %w", err)
	}

	// Create command
	cmd := agentcommand.DeregisterAgentCommand{
		AgentID: agentID,
	}

	// Execute command
	if err := h.deregisterCommandHandler.Handle(ctx, cmd); err != nil {
		return nil, err
	}

	// Convert to response
	return mappers.ToDeregisterAgentResponse(), nil
}

// GetAgent handles getting a single agent by ID
func (h *AgentHandler) GetAgent(ctx context.Context, input *queries.GetAgentRequest) (*queries.GetAgentResponse, error) {
	// Parse agent ID
	agentID, err := agent.NewAgentID(input.ID)
	if err != nil {
		return nil, fmt.Errorf("invalid agent ID: %w", err)
	}

	// Create query
	query := agentquery.GetAgentQuery{
		ID: agentID,
	}

	// Execute query
	a, err := h.getQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	if a == nil {
		return nil, agent.ErrAgentNotFound
	}

	// Convert to response
	return mappers.ToGetAgentResponse(a), nil
}

// ListAgents handles listing agents with optional filters
func (h *AgentHandler) ListAgents(ctx context.Context, input *queries.ListAgentsRequest) (*queries.ListAgentsResponse, error) {
	// Create query
	query := agentquery.ListAgentsQuery{
		HeartbeatTimeout: 5 * time.Minute, // Default timeout
		OnlyHealthy:      input.OnlyHealthy,
	}

	// Convert empty strings to nil pointers for optional filters
	if input.Status != "" {
		query.Status = &input.Status
	}
	if input.BrowserType != "" {
		query.BrowserType = &input.BrowserType
	}
	if input.Platform != "" {
		query.Platform = &input.Platform
	}

	// Parse tag IDs if provided
	if input.TagIDs != "" {
		tagIDStrings := strings.Split(input.TagIDs, ",")
		query.TagIDs = make([]tag.TagID, len(tagIDStrings))
		for i, tagIDStr := range tagIDStrings {
			tagID, err := tag.NewTagID(strings.TrimSpace(tagIDStr))
			if err != nil {
				return nil, fmt.Errorf("invalid tag ID: %w", err)
			}
			query.TagIDs[i] = tagID
		}
	}

	// Execute query
	agents, err := h.listQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	// Convert to response
	return mappers.ToListAgentsResponse(agents), nil
}

// GetAvailableAgents handles getting available agents with optional filters
func (h *AgentHandler) GetAvailableAgents(ctx context.Context, input *queries.GetAvailableAgentsRequest) (*queries.GetAvailableAgentsResponse, error) {
	// Create query
	query := agentquery.GetAvailableAgentsQuery{}

	// Convert empty strings to nil pointers for optional filters
	if input.BrowserType != "" {
		query.BrowserType = &input.BrowserType
	}
	if input.Platform != "" {
		query.Platform = &input.Platform
	}

	// Parse tag IDs if provided
	if input.TagIDs != "" {
		tagIDStrings := strings.Split(input.TagIDs, ",")
		query.TagIDs = make([]tag.TagID, len(tagIDStrings))
		for i, tagIDStr := range tagIDStrings {
			tagID, err := tag.NewTagID(strings.TrimSpace(tagIDStr))
			if err != nil {
				return nil, fmt.Errorf("invalid tag ID: %w", err)
			}
			query.TagIDs[i] = tagID
		}
	}

	// Execute query
	agents, err := h.getAvailableQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	// Convert to response
	return mappers.ToGetAvailableAgentsResponse(agents), nil
}

// GetStaleAgents handles getting stale agents
func (h *AgentHandler) GetStaleAgents(ctx context.Context, input *queries.GetStaleAgentsRequest) (*queries.GetStaleAgentsResponse, error) {
	// Determine heartbeat timeout (use provided value or default to 5)
	timeoutMinutes := input.HeartbeatTimeoutMinutes
	if timeoutMinutes == 0 {
		timeoutMinutes = 5 // default
	}

	// Create query
	query := agentquery.GetStaleAgentsQuery{
		HeartbeatTimeout: time.Duration(timeoutMinutes) * time.Minute,
	}

	// Execute query
	agents, err := h.getStaleQueryHandler.Handle(ctx, query)
	if err != nil {
		return nil, err
	}

	// Convert to response
	return mappers.ToGetStaleAgentsResponse(agents), nil
}
