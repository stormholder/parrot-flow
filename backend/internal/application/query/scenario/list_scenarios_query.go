package query

import (
	"context"
	"parrotflow/internal/domain/scenario"
)

type ListScenariosQuery struct {
	Criteria scenario.SearchCriteria
}

type ListScenariosQueryHandler struct {
	repository scenario.Repository
}

func NewListScenariosQueryHandler(repository scenario.Repository) *ListScenariosQueryHandler {
	return &ListScenariosQueryHandler{
		repository: repository,
	}
}

func (h *ListScenariosQueryHandler) Handle(ctx context.Context, query ListScenariosQuery) ([]*scenario.Scenario, error) {
	return h.repository.FindAll(ctx, query.Criteria)
}
