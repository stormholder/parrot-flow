package query

import (
	"context"
	"parrotflow/internal/domain/run"
)

type ListRunsQuery struct {
	Criteria run.SearchCriteria
}

type ListRunsQueryHandler struct {
	repository run.Repository
}

func NewListRunsQueryHandler(repository run.Repository) *ListRunsQueryHandler {
	return &ListRunsQueryHandler{
		repository: repository,
	}
}

func (h *ListRunsQueryHandler) Handle(ctx context.Context, query ListRunsQuery) ([]*run.Run, error) {
	return h.repository.FindAll(ctx, query.Criteria)
}
