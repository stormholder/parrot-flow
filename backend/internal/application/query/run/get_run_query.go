package query

import (
	"context"
	"parrotflow/internal/domain/run"
)

type GetRunQuery struct {
	ID run.RunID
}

type GetRunQueryHandler struct {
	repository run.Repository
}

func NewGetRunQueryHandler(repository run.Repository) *GetRunQueryHandler {
	return &GetRunQueryHandler{
		repository: repository,
	}
}

func (h *GetRunQueryHandler) Handle(ctx context.Context, query GetRunQuery) (*run.Run, error) {
	return h.repository.FindByID(ctx, query.ID)
}
