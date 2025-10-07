package query

import (
	"context"
	"parrotflow/internal/domain/scenario"
)

type GetScenarioQuery struct {
	ID scenario.ScenarioID
}

type GetScenarioQueryHandler struct {
	repository scenario.Repository
}

func NewGetScenarioQueryHandler(repository scenario.Repository) *GetScenarioQueryHandler {
	return &GetScenarioQueryHandler{
		repository: repository,
	}
}

func (h *GetScenarioQueryHandler) Handle(ctx context.Context, query GetScenarioQuery) (*scenario.Scenario, error) {
	return h.repository.FindByID(ctx, query.ID)
}
