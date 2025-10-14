package proxy

import (
	"context"

	"parrotflow/internal/domain/proxy"
)

type ListProxiesQuery struct {
	Status *string   // Optional filter by status
	Tags   []string  // Optional filter by tags
}

type ListProxiesQueryHandler struct {
	repository proxy.ProxyRepository
}

func NewListProxiesQueryHandler(repository proxy.ProxyRepository) *ListProxiesQueryHandler {
	return &ListProxiesQueryHandler{
		repository: repository,
	}
}

func (h *ListProxiesQueryHandler) Handle(ctx context.Context, query ListProxiesQuery) ([]*proxy.Proxy, error) {
	// If filtering by tags
	if len(query.Tags) > 0 {
		return h.repository.FindByTags(ctx, query.Tags)
	}

	// If filtering by status
	if query.Status != nil {
		status, err := proxy.NewProxyStatus(*query.Status)
		if err != nil {
			return nil, err
		}
		return h.repository.FindByStatus(ctx, status)
	}

	// Return all proxies
	// Note: In production, this should support pagination
	// For now, we'll use a simple approach - return all via FindActive
	// which is safer than returning truly all records
	return h.repository.FindActive(ctx)
}
