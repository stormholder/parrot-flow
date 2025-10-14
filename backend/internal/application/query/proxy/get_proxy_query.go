package proxy

import (
	"context"

	"parrotflow/internal/domain/proxy"
)

type GetProxyQuery struct {
	ID proxy.ProxyID
}

type GetProxyQueryHandler struct {
	repository proxy.ProxyRepository
}

func NewGetProxyQueryHandler(repository proxy.ProxyRepository) *GetProxyQueryHandler {
	return &GetProxyQueryHandler{
		repository: repository,
	}
}

func (h *GetProxyQueryHandler) Handle(ctx context.Context, query GetProxyQuery) (*proxy.Proxy, error) {
	p, err := h.repository.FindByID(ctx, query.ID)
	if err != nil {
		return nil, err
	}
	if p == nil {
		return nil, proxy.ErrProxyNotFound
	}
	return p, nil
}
