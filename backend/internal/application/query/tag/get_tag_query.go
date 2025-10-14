package query

import (
	"context"
	"parrotflow/internal/domain/tag"
)

type GetTagQuery struct {
	ID tag.TagID
}

type GetTagQueryHandler struct {
	repository tag.Repository
}

func NewGetTagQueryHandler(repository tag.Repository) *GetTagQueryHandler {
	return &GetTagQueryHandler{
		repository: repository,
	}
}

func (h *GetTagQueryHandler) Handle(ctx context.Context, query GetTagQuery) (*tag.Tag, error) {
	return h.repository.FindByID(ctx, query.ID)
}
