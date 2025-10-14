package mappers

import (
	"parrotflow/internal/domain/proxy"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/queries"
)

// ToCreateProxyResponse converts a domain Proxy to CreateProxyResponse
func ToCreateProxyResponse(p *proxy.Proxy) *commands.CreateProxyResponse {
	response := &commands.CreateProxyResponse{}
	response.Body.ID = p.Id.String()
	response.Body.Name = p.Name
	response.Body.Host = p.Host
	response.Body.Port = p.Port
	response.Body.Protocol = p.Protocol.String()
	response.Body.Status = p.Status.String()
	response.Body.HasCredentials = p.Credentials != nil
	response.Body.FailureCount = p.FailureCount
	response.Body.SuccessCount = p.SuccessCount
	response.Body.AverageLatency = p.AverageLatency
	response.Body.CreatedAt = p.CreatedAt.Time().Format("2006-01-02T15:04:05Z07:00")
	response.Body.UpdatedAt = p.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00")

	// Convert tags
	response.Body.Tags = make([]string, len(p.Tags))
	for i, tagID := range p.Tags {
		response.Body.Tags[i] = tagID.String()
	}

	// Convert LastCheckedAt if present
	if p.LastCheckedAt != nil {
		formatted := p.LastCheckedAt.Time().Format("2006-01-02T15:04:05Z07:00")
		response.Body.LastCheckedAt = &formatted
	}

	return response
}

// ToUpdateProxyResponse converts a domain Proxy to UpdateProxyResponse
func ToUpdateProxyResponse(p *proxy.Proxy) *commands.UpdateProxyResponse {
	response := &commands.UpdateProxyResponse{}
	response.Body.ID = p.Id.String()
	response.Body.Name = p.Name
	response.Body.Host = p.Host
	response.Body.Port = p.Port
	response.Body.Protocol = p.Protocol.String()
	response.Body.Status = p.Status.String()
	response.Body.HasCredentials = p.Credentials != nil
	response.Body.FailureCount = p.FailureCount
	response.Body.SuccessCount = p.SuccessCount
	response.Body.AverageLatency = p.AverageLatency
	response.Body.CreatedAt = p.CreatedAt.Time().Format("2006-01-02T15:04:05Z07:00")
	response.Body.UpdatedAt = p.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00")

	// Convert tags
	response.Body.Tags = make([]string, len(p.Tags))
	for i, tagID := range p.Tags {
		response.Body.Tags[i] = tagID.String()
	}

	// Convert LastCheckedAt if present
	if p.LastCheckedAt != nil {
		formatted := p.LastCheckedAt.Time().Format("2006-01-02T15:04:05Z07:00")
		response.Body.LastCheckedAt = &formatted
	}

	return response
}

// ToRecordHealthResponse converts a domain Proxy to RecordHealthResponse
func ToRecordHealthResponse(p *proxy.Proxy) *commands.RecordHealthResponse {
	response := &commands.RecordHealthResponse{}
	response.Body.ID = p.Id.String()
	response.Body.Name = p.Name
	response.Body.Host = p.Host
	response.Body.Port = p.Port
	response.Body.Protocol = p.Protocol.String()
	response.Body.Status = p.Status.String()
	response.Body.HasCredentials = p.Credentials != nil
	response.Body.FailureCount = p.FailureCount
	response.Body.SuccessCount = p.SuccessCount
	response.Body.AverageLatency = p.AverageLatency
	response.Body.CreatedAt = p.CreatedAt.Time().Format("2006-01-02T15:04:05Z07:00")
	response.Body.UpdatedAt = p.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00")

	// Convert tags
	response.Body.Tags = make([]string, len(p.Tags))
	for i, tagID := range p.Tags {
		response.Body.Tags[i] = tagID.String()
	}

	// Convert LastCheckedAt if present
	if p.LastCheckedAt != nil {
		formatted := p.LastCheckedAt.Time().Format("2006-01-02T15:04:05Z07:00")
		response.Body.LastCheckedAt = &formatted
	}

	return response
}

// ToActivateProxyResponse converts a domain Proxy to ActivateProxyResponse
func ToActivateProxyResponse(p *proxy.Proxy) *commands.ActivateProxyResponse {
	response := &commands.ActivateProxyResponse{}
	response.Body.ID = p.Id.String()
	response.Body.Name = p.Name
	response.Body.Status = p.Status.String()
	return response
}

// ToDeactivateProxyResponse converts a domain Proxy to DeactivateProxyResponse
func ToDeactivateProxyResponse(p *proxy.Proxy) *commands.DeactivateProxyResponse {
	response := &commands.DeactivateProxyResponse{}
	response.Body.ID = p.Id.String()
	response.Body.Name = p.Name
	response.Body.Status = p.Status.String()
	return response
}

// ToGetProxyResponse converts a domain Proxy to GetProxyResponse
func ToGetProxyResponse(p *proxy.Proxy) *queries.GetProxyResponse {
	response := &queries.GetProxyResponse{}
	response.Body = toProxyDTO(p)
	return response
}

// ToListProxiesResponse converts a slice of domain Proxies to ListProxiesResponse
func ToListProxiesResponse(proxies []*proxy.Proxy) *queries.ListProxiesResponse {
	response := &queries.ListProxiesResponse{}
	response.Body.Proxies = make([]queries.ProxyDTO, len(proxies))
	for i, p := range proxies {
		response.Body.Proxies[i] = toProxyDTO(p)
	}
	response.Body.Total = len(proxies)
	return response
}

// ToGetActiveProxiesResponse converts a slice of domain Proxies to GetActiveProxiesResponse
func ToGetActiveProxiesResponse(proxies []*proxy.Proxy) *queries.GetActiveProxiesResponse {
	response := &queries.GetActiveProxiesResponse{}
	response.Body.Proxies = make([]queries.ProxyDTO, len(proxies))
	for i, p := range proxies {
		response.Body.Proxies[i] = toProxyDTO(p)
	}
	response.Body.Total = len(proxies)
	return response
}

// toProxyDTO converts a domain Proxy to ProxyDTO (helper function)
func toProxyDTO(p *proxy.Proxy) queries.ProxyDTO {
	dto := queries.ProxyDTO{
		ID:             p.Id.String(),
		Name:           p.Name,
		Host:           p.Host,
		Port:           p.Port,
		Protocol:       p.Protocol.String(),
		Status:         p.Status.String(),
		ConnectionURL:  p.GetConnectionURL(),
		HasCredentials: p.Credentials != nil,
		FailureCount:   p.FailureCount,
		SuccessCount:   p.SuccessCount,
		AverageLatency: p.AverageLatency,
		CreatedAt:      p.CreatedAt.Time().Format("2006-01-02T15:04:05Z07:00"),
		UpdatedAt:      p.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00"),
	}

	// Convert tags
	dto.Tags = make([]string, len(p.Tags))
	for i, tagID := range p.Tags {
		dto.Tags[i] = tagID.String()
	}

	// Convert LastCheckedAt if present
	if p.LastCheckedAt != nil {
		formatted := p.LastCheckedAt.Time().Format("2006-01-02T15:04:05Z07:00")
		dto.LastCheckedAt = &formatted
	}

	// Convert LastFailureAt if present
	if p.LastFailureAt != nil {
		formatted := p.LastFailureAt.Time().Format("2006-01-02T15:04:05Z07:00")
		dto.LastFailureAt = &formatted
	}

	return dto
}
