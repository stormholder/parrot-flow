package query

import (
	"context"
	"parrotflow/internal/domain/tag"
)

type ListTagsQuery struct {
	Category *tag.TagCategory // Optional filter by category
}

type ListTagsQueryHandler struct {
	repository tag.Repository
}

func NewListTagsQueryHandler(repository tag.Repository) *ListTagsQueryHandler {
	return &ListTagsQueryHandler{
		repository: repository,
	}
}

func (h *ListTagsQueryHandler) Handle(ctx context.Context, query ListTagsQuery) ([]*tag.Tag, error) {
	if query.Category != nil {
		return h.repository.FindByCategory(ctx, *query.Category)
	}
	return h.repository.FindAll(ctx)
}
