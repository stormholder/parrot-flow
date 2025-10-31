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

// Mapper functions using functional approach

func ProxyToCreateResponse(p *proxy.Proxy) *commands.CreateProxyResponse {
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

func ProxyToUpdateResponse(p *proxy.Proxy) *commands.UpdateProxyResponse {
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

func ProxyToDeleteResponse() *commands.DeleteProxyResponse {
	response := &commands.DeleteProxyResponse{}
	response.Body.Message = "Proxy deleted successfully"
	return response
}

func ProxyToActivateResponse(p *proxy.Proxy) *commands.ActivateProxyResponse {
	response := &commands.ActivateProxyResponse{}
	response.Body.ID = p.Id.String()
	response.Body.Status = p.Status.String()
	return response
}

func ProxyToDeactivateResponse(p *proxy.Proxy) *commands.DeactivateProxyResponse {
	response := &commands.DeactivateProxyResponse{}
	response.Body.ID = p.Id.String()
	response.Body.Status = p.Status.String()
	return response
}

func ProxyToRecordHealthResponse(p *proxy.Proxy) *commands.RecordHealthResponse {
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

func ProxyToGetResponse(p *proxy.Proxy) *queries.GetProxyResponse {
	response := &queries.GetProxyResponse{}
	response.Body = buildProxyDTO(p)
	return response
}

func ProxyToListResponse(proxies []*proxy.Proxy) *queries.ListProxiesResponse {
	response := &queries.ListProxiesResponse{}
	response.Body.Proxies = MapSlicePtr(proxies, buildProxyDTO)
	return response
}

func ProxyToActiveListResponse(proxies []*proxy.Proxy) *queries.GetActiveProxiesResponse {
	response := &queries.GetActiveProxiesResponse{}
	response.Body.Proxies = MapSlicePtr(proxies, buildProxyDTO)
	return response
}

// Mapper instances for handler injection - using functional types
var (
	ProxyCreateMapper      = CreateMapperFunc[*proxy.Proxy, *commands.CreateProxyResponse](ProxyToCreateResponse)
	ProxyUpdateMapper      = UpdateMapperFunc[*proxy.Proxy, *commands.UpdateProxyResponse](ProxyToUpdateResponse)
	ProxyDeleteMapper      = DeleteMapperFunc[*commands.DeleteProxyResponse](ProxyToDeleteResponse)
	ProxyActivateMapper    = CreateMapperFunc[*proxy.Proxy, *commands.ActivateProxyResponse](ProxyToActivateResponse)
	ProxyDeactivateMapper  = CreateMapperFunc[*proxy.Proxy, *commands.DeactivateProxyResponse](ProxyToDeactivateResponse)
	ProxyRecordHealthMapper = CreateMapperFunc[*proxy.Proxy, *commands.RecordHealthResponse](ProxyToRecordHealthResponse)
	ProxyGetMapper         = GetMapperFunc[*proxy.Proxy, *queries.GetProxyResponse](ProxyToGetResponse)
	ProxyListMapper        = ListMapperFunc[proxy.Proxy, *queries.ListProxiesResponse](ProxyToListResponse)
	ProxyActiveListMapper  = ListMapperFunc[proxy.Proxy, *queries.GetActiveProxiesResponse](ProxyToActiveListResponse)
)
