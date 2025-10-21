package handlers

import (
	"context"
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
	// Commands
	registerCommandHandler       *agentcommand.RegisterAgentCommandHandler
	updateHeartbeatCommandHandler *agentcommand.UpdateHeartbeatCommandHandler
	assignRunCommandHandler      *agentcommand.AssignRunCommandHandler
	releaseRunCommandHandler     *agentcommand.ReleaseRunCommandHandler
	updateCommandHandler         *agentcommand.UpdateAgentCommandHandler
	deregisterCommandHandler     *agentcommand.DeregisterAgentCommandHandler

	// Queries
	getQueryHandler          *agentquery.GetAgentQueryHandler
	listQueryHandler         *agentquery.ListAgentsQueryHandler
	getAvailableQueryHandler *agentquery.GetAvailableAgentsQueryHandler
	getStaleQueryHandler     *agentquery.GetStaleAgentsQueryHandler

	// Mappers
	registerMapper       mappers.AgentRegisterMapper
	heartbeatMapper      mappers.AgentHeartbeatMapper
	assignRunMapper      mappers.AgentAssignRunMapper
	releaseRunMapper     mappers.AgentReleaseRunMapper
	updateMapper         mappers.AgentUpdateMapper
	deregisterMapper     mappers.AgentDeregisterMapper
	getMapper            mappers.AgentGetMapper
	listMapper           mappers.AgentListMapper
	availableListMapper  mappers.AgentAvailableListMapper
	staleListMapper      mappers.AgentStaleListMapper
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
		registerCommandHandler:       registerCommandHandler,
		updateHeartbeatCommandHandler: updateHeartbeatCommandHandler,
		assignRunCommandHandler:      assignRunCommandHandler,
		releaseRunCommandHandler:     releaseRunCommandHandler,
		updateCommandHandler:         updateCommandHandler,
		deregisterCommandHandler:     deregisterCommandHandler,
		getQueryHandler:              getQueryHandler,
		listQueryHandler:             listQueryHandler,
		getAvailableQueryHandler:     getAvailableQueryHandler,
		getStaleQueryHandler:         getStaleQueryHandler,
		registerMapper:               mappers.AgentRegisterMapper{},
		heartbeatMapper:              mappers.AgentHeartbeatMapper{},
		assignRunMapper:              mappers.AgentAssignRunMapper{},
		releaseRunMapper:             mappers.AgentReleaseRunMapper{},
		updateMapper:                 mappers.AgentUpdateMapper{},
		deregisterMapper:             mappers.AgentDeregisterMapper{},
		getMapper:                    mappers.AgentGetMapper{},
		listMapper:                   mappers.AgentListMapper{},
		availableListMapper:          mappers.AgentAvailableListMapper{},
		staleListMapper:              mappers.AgentStaleListMapper{},
	}
}

func (h *AgentHandler) RegisterAgent(ctx context.Context, req *commands.RegisterAgentRequest) (*commands.RegisterAgentResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.RegisterAgentRequest) (agentcommand.RegisterAgentCommand, error) {
			capabilities, err := mappers.FromCapabilitiesDTO(r.Body.Capabilities)
			if err != nil {
				return agentcommand.RegisterAgentCommand{}, err
			}

			connectionInfo, err := mappers.FromConnectionInfoDTO(r.Body.ConnectionInfo)
			if err != nil {
				return agentcommand.RegisterAgentCommand{}, err
			}

			return agentcommand.RegisterAgentCommand{
				Name:           r.Body.Name,
				Capabilities:   capabilities,
				ConnectionInfo: connectionInfo,
			}, nil
		},
		CommandHandlerFunc[agentcommand.RegisterAgentCommand, *agent.Agent](h.registerCommandHandler.Handle),
		h.registerMapper,
	)
}

func (h *AgentHandler) UpdateHeartbeat(ctx context.Context, req *commands.UpdateHeartbeatRequest) (*commands.UpdateHeartbeatResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.UpdateHeartbeatRequest) (agentcommand.UpdateHeartbeatCommand, error) {
			agentID, err := agent.NewAgentID(r.ID)
			if err != nil {
				return agentcommand.UpdateHeartbeatCommand{}, err
			}
			return agentcommand.UpdateHeartbeatCommand{AgentID: agentID}, nil
		},
		CommandHandlerFunc[agentcommand.UpdateHeartbeatCommand, *agent.Agent](h.updateHeartbeatCommandHandler.Handle),
		h.heartbeatMapper,
	)
}

func (h *AgentHandler) AssignRun(ctx context.Context, req *commands.AssignRunRequest) (*commands.AssignRunResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.AssignRunRequest) (agentcommand.AssignRunCommand, error) {
			agentID, err := agent.NewAgentID(r.ID)
			if err != nil {
				return agentcommand.AssignRunCommand{}, err
			}
			return agentcommand.AssignRunCommand{AgentID: agentID}, nil
		},
		CommandHandlerFunc[agentcommand.AssignRunCommand, *agent.Agent](h.assignRunCommandHandler.Handle),
		h.assignRunMapper,
	)
}

func (h *AgentHandler) ReleaseRun(ctx context.Context, req *commands.ReleaseRunRequest) (*commands.ReleaseRunResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.ReleaseRunRequest) (agentcommand.ReleaseRunCommand, error) {
			agentID, err := agent.NewAgentID(r.ID)
			if err != nil {
				return agentcommand.ReleaseRunCommand{}, err
			}
			return agentcommand.ReleaseRunCommand{AgentID: agentID}, nil
		},
		CommandHandlerFunc[agentcommand.ReleaseRunCommand, *agent.Agent](h.releaseRunCommandHandler.Handle),
		h.releaseRunMapper,
	)
}

func (h *AgentHandler) UpdateAgent(ctx context.Context, req *commands.UpdateAgentRequest) (*commands.UpdateAgentResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.UpdateAgentRequest) (agentcommand.UpdateAgentCommand, error) {
			agentID, err := agent.NewAgentID(r.ID)
			if err != nil {
				return agentcommand.UpdateAgentCommand{}, err
			}

			var capabilities *agent.Capabilities
			if r.Body.Capabilities != nil {
				caps, err := mappers.FromCapabilitiesDTO(*r.Body.Capabilities)
				if err != nil {
					return agentcommand.UpdateAgentCommand{}, err
				}
				capabilities = &caps
			}

			var tagsToAdd []tag.TagID
			if len(r.Body.TagsToAdd) > 0 {
				tagsToAdd = make([]tag.TagID, len(r.Body.TagsToAdd))
				for i, idStr := range r.Body.TagsToAdd {
					tagID, err := tag.NewTagID(idStr)
					if err != nil {
						return agentcommand.UpdateAgentCommand{}, err
					}
					tagsToAdd[i] = tagID
				}
			}

			var tagsToRemove []tag.TagID
			if len(r.Body.TagsToRemove) > 0 {
				tagsToRemove = make([]tag.TagID, len(r.Body.TagsToRemove))
				for i, idStr := range r.Body.TagsToRemove {
					tagID, err := tag.NewTagID(idStr)
					if err != nil {
						return agentcommand.UpdateAgentCommand{}, err
					}
					tagsToRemove[i] = tagID
				}
			}

			return agentcommand.UpdateAgentCommand{
				AgentID:      agentID,
				Name:         r.Body.Name,
				Capabilities: capabilities,
				TagsToAdd:    tagsToAdd,
				TagsToRemove: tagsToRemove,
			}, nil
		},
		CommandHandlerFunc[agentcommand.UpdateAgentCommand, *agent.Agent](h.updateCommandHandler.Handle),
		h.updateMapper,
	)
}

func (h *AgentHandler) DeregisterAgent(ctx context.Context, req *commands.DeregisterAgentRequest) (*commands.DeregisterAgentResponse, error) {
	return HandleSimpleCommand(
		ctx,
		req,
		func(r *commands.DeregisterAgentRequest) (agentcommand.DeregisterAgentCommand, error) {
			agentID, err := agent.NewAgentID(r.ID)
			if err != nil {
				return agentcommand.DeregisterAgentCommand{}, err
			}
			return agentcommand.DeregisterAgentCommand{AgentID: agentID}, nil
		},
		SimpleCommandHandlerFunc[agentcommand.DeregisterAgentCommand](h.deregisterCommandHandler.Handle),
		h.deregisterMapper.Map,
	)
}

func (h *AgentHandler) GetAgent(ctx context.Context, req *queries.GetAgentRequest) (*queries.GetAgentResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.GetAgentRequest) (agentquery.GetAgentQuery, error) {
			agentID, err := agent.NewAgentID(r.ID)
			if err != nil {
				return agentquery.GetAgentQuery{}, err
			}
			return agentquery.GetAgentQuery{ID: agentID}, nil
		},
		QueryHandlerFunc[agentquery.GetAgentQuery, *agent.Agent](h.getQueryHandler.Handle),
		h.getMapper,
	)
}

func (h *AgentHandler) ListAgents(ctx context.Context, req *queries.ListAgentsRequest) (*queries.ListAgentsResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.ListAgentsRequest) (agentquery.ListAgentsQuery, error) {
			q := agentquery.ListAgentsQuery{}

			if r.Status != "" {
				q.Status = &r.Status
			}

			// Note: TagIDs is a string in the request, needs parsing if we want to use it
			// For now, we'll leave it empty
			return q, nil
		},
		QueryHandlerFunc[agentquery.ListAgentsQuery, []*agent.Agent](h.listQueryHandler.Handle),
		h.listMapper,
	)
}

func (h *AgentHandler) GetAvailableAgents(ctx context.Context, req *queries.GetAvailableAgentsRequest) (*queries.GetAvailableAgentsResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.GetAvailableAgentsRequest) (agentquery.GetAvailableAgentsQuery, error) {
			return agentquery.GetAvailableAgentsQuery{}, nil
		},
		QueryHandlerFunc[agentquery.GetAvailableAgentsQuery, []*agent.Agent](h.getAvailableQueryHandler.Handle),
		h.availableListMapper,
	)
}

func (h *AgentHandler) GetStaleAgents(ctx context.Context, req *queries.GetStaleAgentsRequest) (*queries.GetStaleAgentsResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.GetStaleAgentsRequest) (agentquery.GetStaleAgentsQuery, error) {
			timeout := 5 * time.Minute
			if r.HeartbeatTimeoutMinutes > 0 {
				timeout = time.Duration(r.HeartbeatTimeoutMinutes) * time.Minute
			}
			return agentquery.GetStaleAgentsQuery{HeartbeatTimeout: timeout}, nil
		},
		QueryHandlerFunc[agentquery.GetStaleAgentsQuery, []*agent.Agent](h.getStaleQueryHandler.Handle),
		h.staleListMapper,
	)
}
