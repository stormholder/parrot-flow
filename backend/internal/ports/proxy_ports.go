package ports

import (
	"parrotflow/internal/domain/proxy"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/models"
)

func ProxyParseID(id string) uint64 {
	return parseID(id)
}

func ProxyDomainEntityToPersistence(p *proxy.Proxy) (*models.Proxy, error) {
	model := &models.Proxy{
		Model: models.Model{
			ID:        parseID(p.Id.String()),
			CreatedAt: p.CreatedAt.Time(),
			UpdatedAt: p.UpdatedAt.Time(),
		},
		Name:           p.Name,
		Host:           p.Host,
		Port:           p.Port,
		Protocol:       p.Protocol.String(),
		Status:         p.Status.String(),
		FailureCount:   p.FailureCount,
		SuccessCount:   p.SuccessCount,
		AverageLatency: p.AverageLatency,
	}

	// Handle credentials
	if p.Credentials != nil {
		model.Username = p.Credentials.Username
		model.Password = p.Credentials.Password // TODO: Encryption in future
	}

	// Handle timestamps
	if p.LastCheckedAt != nil {
		t := p.LastCheckedAt.Time()
		model.LastCheckedAt = &t
	}
	if p.LastFailureAt != nil {
		t := p.LastFailureAt.Time()
		model.LastFailureAt = &t
	}

	// Handle tags - we'll load them separately in repository
	// Tags are stored as IDs in domain, need to be resolved to full Tag models

	return model, nil
}

func ProxyPersistenceToDomainEntity(model *models.Proxy) (*proxy.Proxy, error) {
	proxyID, err := proxy.NewProxyID(formatID(model.ID))
	if err != nil {
		return nil, err
	}

	protocol, err := proxy.NewProxyProtocol(model.Protocol)
	if err != nil {
		return nil, err
	}

	p, err := proxy.NewProxy(proxyID, model.Name, model.Host, model.Port, protocol)
	if err != nil {
		return nil, err
	}

	// Set credentials if present
	if model.Username != "" && model.Password != "" {
		creds, err := proxy.NewProxyCredentials(model.Username, model.Password)
		if err != nil {
			return nil, err
		}
		p.SetCredentials(creds)
	}

	// Set status
	status, err := proxy.NewProxyStatus(model.Status)
	if err != nil {
		return nil, err
	}
	p.UpdateStatus(status)

	// Set metrics
	p.FailureCount = model.FailureCount
	p.SuccessCount = model.SuccessCount
	p.AverageLatency = model.AverageLatency

	// Set timestamps
	if model.LastCheckedAt != nil {
		ts := shared.NewTimestamp(*model.LastCheckedAt)
		p.LastCheckedAt = &ts
	}
	if model.LastFailureAt != nil {
		ts := shared.NewTimestamp(*model.LastFailureAt)
		p.LastFailureAt = &ts
	}

	// Convert tags
	for _, tagModel := range model.Tags {
		tagID, err := tag.NewTagID(formatID(tagModel.ID))
		if err != nil {
			return nil, err
		}
		p.AddTag(tagID)
	}

	return p, nil
}
