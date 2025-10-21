package mappers

import (
	"parrotflow/internal/domain/proxy"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/queries"
)

func buildProxyDTO(p *proxy.Proxy) queries.ProxyDTO {
	var lastChecked *string
	if p.LastCheckedAt != nil {
		s := FormatTimestamp(p.LastCheckedAt.Time())
		lastChecked = &s
	}

	var lastFailure *string
	if p.LastFailureAt != nil {
		s := FormatTimestamp(p.LastFailureAt.Time())
		lastFailure = &s
	}

	tagStrings := make([]string, len(p.Tags))
	for i, tagID := range p.Tags {
		tagStrings[i] = tagID.String()
	}

	return queries.ProxyDTO{
		ID:             p.Id.String(),
		Name:           p.Name,
		Host:           p.Host,
		Port:           p.Port,
		Protocol:       p.Protocol.String(),
		Status:         p.Status.String(),
		ConnectionURL:  p.GetConnectionURL(),
		HasCredentials: p.Credentials != nil,
		Tags:           tagStrings,
		LastCheckedAt:  lastChecked,
		LastFailureAt:  lastFailure,
		FailureCount:   p.FailureCount,
		SuccessCount:   p.SuccessCount,
		AverageLatency: p.AverageLatency,
		CreatedAt:      FormatTimestamp(p.CreatedAt.Time()),
	}
}

type ProxyCreateMapper struct{}

func (m ProxyCreateMapper) Map(p *proxy.Proxy) *commands.CreateProxyResponse {
	dto := buildProxyDTO(p)
	response := &commands.CreateProxyResponse{}
	response.Body.ID = dto.ID
	response.Body.Name = dto.Name
	response.Body.Host = dto.Host
	response.Body.Port = dto.Port
	response.Body.Protocol = dto.Protocol
	response.Body.Status = dto.Status
	return response
}

type ProxyUpdateMapper struct{}

func (m ProxyUpdateMapper) Map(p *proxy.Proxy) *commands.UpdateProxyResponse {
	dto := buildProxyDTO(p)
	response := &commands.UpdateProxyResponse{}
	response.Body.ID = dto.ID
	response.Body.Name = dto.Name
	response.Body.Host = dto.Host
	response.Body.Port = dto.Port
	response.Body.Protocol = dto.Protocol
	response.Body.Status = dto.Status
	response.Body.Tags = dto.Tags
	return response
}

type ProxyDeleteMapper struct{}

func (m ProxyDeleteMapper) Map() *commands.DeleteProxyResponse {
	response := &commands.DeleteProxyResponse{}
	response.Body.Message = "Proxy deleted successfully"
	return response
}

type ProxyActivateMapper struct{}

func (m ProxyActivateMapper) Map(p *proxy.Proxy) *commands.ActivateProxyResponse {
	response := &commands.ActivateProxyResponse{}
	response.Body.ID = p.Id.String()
	response.Body.Status = p.Status.String()
	return response
}

type ProxyDeactivateMapper struct{}

func (m ProxyDeactivateMapper) Map(p *proxy.Proxy) *commands.DeactivateProxyResponse {
	response := &commands.DeactivateProxyResponse{}
	response.Body.ID = p.Id.String()
	response.Body.Status = p.Status.String()
	return response
}

type ProxyRecordHealthMapper struct{}

func (m ProxyRecordHealthMapper) Map(p *proxy.Proxy) *commands.RecordHealthResponse {
	dto := buildProxyDTO(p)
	response := &commands.RecordHealthResponse{}
	response.Body.ID = dto.ID
	response.Body.Status = dto.Status
	response.Body.FailureCount = dto.FailureCount
	response.Body.SuccessCount = dto.SuccessCount
	response.Body.AverageLatency = dto.AverageLatency
	response.Body.LastCheckedAt = dto.LastCheckedAt
	return response
}

type ProxyGetMapper struct{}

func (m ProxyGetMapper) Map(p *proxy.Proxy) *queries.GetProxyResponse {
	response := &queries.GetProxyResponse{}
	response.Body = buildProxyDTO(p)
	return response
}

type ProxyListMapper struct{}

func (m ProxyListMapper) Map(proxies []*proxy.Proxy) *queries.ListProxiesResponse {
	response := &queries.ListProxiesResponse{}
	response.Body.Proxies = MapSlicePtr(proxies, buildProxyDTO)
	return response
}

type ProxyActiveListMapper struct{}

func (m ProxyActiveListMapper) Map(proxies []*proxy.Proxy) *queries.GetActiveProxiesResponse {
	response := &queries.GetActiveProxiesResponse{}
	response.Body.Proxies = MapSlicePtr(proxies, buildProxyDTO)
	return response
}

func ToCreateProxyResponse(p *proxy.Proxy) *commands.CreateProxyResponse {
	return ProxyCreateMapper{}.Map(p)
}

func ToUpdateProxyResponse(p *proxy.Proxy) *commands.UpdateProxyResponse {
	return ProxyUpdateMapper{}.Map(p)
}

func ToActivateProxyResponse(p *proxy.Proxy) *commands.ActivateProxyResponse {
	return ProxyActivateMapper{}.Map(p)
}

func ToDeactivateProxyResponse(p *proxy.Proxy) *commands.DeactivateProxyResponse {
	return ProxyDeactivateMapper{}.Map(p)
}

func ToRecordHealthResponse(p *proxy.Proxy) *commands.RecordHealthResponse {
	return ProxyRecordHealthMapper{}.Map(p)
}

func ToGetProxyResponse(p *proxy.Proxy) *queries.GetProxyResponse {
	return ProxyGetMapper{}.Map(p)
}

func ToListProxiesResponse(proxies []*proxy.Proxy) *queries.ListProxiesResponse {
	return ProxyListMapper{}.Map(proxies)
}

func ToGetActiveProxiesResponse(proxies []*proxy.Proxy) *queries.GetActiveProxiesResponse {
	return ProxyActiveListMapper{}.Map(proxies)
}
