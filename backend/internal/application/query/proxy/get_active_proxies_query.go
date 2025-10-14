package proxy

import (
	"context"

	"parrotflow/internal/domain/proxy"
)

type GetActiveProxiesQuery struct {
	// No filter parameters - just get all active proxies
}

type GetActiveProxiesQueryHandler struct {
	repository proxy.Repository
}

func NewGetActiveProxiesQueryHandler(repository proxy.Repository) *GetActiveProxiesQueryHandler {
	return &GetActiveProxiesQueryHandler{
		repository: repository,
	}
}

func (h *GetActiveProxiesQueryHandler) Handle(ctx context.Context, query GetActiveProxiesQuery) ([]*proxy.Proxy, error) {
	return h.repository.FindActive(ctx)
}
