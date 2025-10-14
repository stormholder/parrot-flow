package proxy

import (
	"errors"
	"fmt"
	"parrotflow/internal/domain/shared"
	"parrotflow/internal/domain/tag"
	"time"
)

// Domain errors
var (
	ErrProxyAlreadyExists = errors.New("proxy already exists")
	ErrProxyNotFound      = errors.New("proxy not found")
)

type ProxyID struct {
	shared.ID
}

func NewProxyID(value string) (ProxyID, error) {
	id, err := shared.NewID(value)
	if err != nil {
		return ProxyID{}, err
	}
	return ProxyID{ID: id}, nil
}

// ProxyProtocol represents the proxy protocol type
type ProxyProtocol struct {
	value string
}

func NewProxyProtocol(value string) (ProxyProtocol, error) {
	switch value {
	case "http", "https", "socks5":
		return ProxyProtocol{value: value}, nil
	default:
		return ProxyProtocol{}, fmt.Errorf("invalid proxy protocol: %s (must be http, https, or socks5)", value)
	}
}

func (pp ProxyProtocol) String() string {
	return pp.value
}

// Common proxy protocols
var (
	ProtocolHTTP   = ProxyProtocol{value: "http"}
	ProtocolHTTPS  = ProxyProtocol{value: "https"}
	ProtocolSOCKS5 = ProxyProtocol{value: "socks5"}
)

// ProxyCredentials represents authentication credentials for a proxy
type ProxyCredentials struct {
	Username string
	Password string // Should be encrypted in storage
}

func NewProxyCredentials(username, password string) (ProxyCredentials, error) {
	if username == "" {
		return ProxyCredentials{}, errors.New("proxy username cannot be empty")
	}
	if password == "" {
		return ProxyCredentials{}, errors.New("proxy password cannot be empty")
	}
	return ProxyCredentials{
		Username: username,
		Password: password,
	}, nil
}

func (pc ProxyCredentials) IsEmpty() bool {
	return pc.Username == "" && pc.Password == ""
}

// ProxyStatus represents the current status of a proxy
type ProxyStatus struct {
	value string
}

func NewProxyStatus(value string) (ProxyStatus, error) {
	switch value {
	case "active", "inactive", "checking", "failed":
		return ProxyStatus{value: value}, nil
	default:
		return ProxyStatus{}, fmt.Errorf("invalid proxy status: %s", value)
	}
}

func (ps ProxyStatus) String() string {
	return ps.value
}

// Common proxy statuses
var (
	ProxyStatusActive   = ProxyStatus{value: "active"}
	ProxyStatusInactive = ProxyStatus{value: "inactive"}
	ProxyStatusChecking = ProxyStatus{value: "checking"}
	ProxyStatusFailed   = ProxyStatus{value: "failed"}
)

// Proxy represents a proxy server that agents can use
type Proxy struct {
	Id              ProxyID
	Name            string
	Host            string
	Port            int
	Protocol        ProxyProtocol
	Credentials     *ProxyCredentials // Nullable - some proxies don't require auth
	Status          ProxyStatus
	Tags            []tag.TagID
	LastCheckedAt   *shared.Timestamp
	LastFailureAt   *shared.Timestamp
	FailureCount    int
	SuccessCount    int
	AverageLatency  int // in milliseconds
	CreatedAt       shared.Timestamp
	UpdatedAt       shared.Timestamp
	Events          []shared.DomainEvent
}

// NewProxy creates a new proxy
func NewProxy(id ProxyID, name, host string, port int, protocol ProxyProtocol) (*Proxy, error) {
	if name == "" {
		return nil, errors.New("proxy name cannot be empty")
	}
	if host == "" {
		return nil, errors.New("proxy host cannot be empty")
	}
	if port <= 0 || port > 65535 {
		return nil, errors.New("proxy port must be between 1 and 65535")
	}

	proxy := &Proxy{
		Id:           id,
		Name:         name,
		Host:         host,
		Port:         port,
		Protocol:     protocol,
		Status:       ProxyStatusInactive,
		Tags:         make([]tag.TagID, 0),
		FailureCount: 0,
		SuccessCount: 0,
		CreatedAt:    shared.NewTimestamp(time.Now()),
		UpdatedAt:    shared.NewTimestamp(time.Now()),
		Events:       make([]shared.DomainEvent, 0),
	}

	proxy.addEvent(ProxyCreated{
		BaseEvent: shared.NewBaseEvent(EventProxyCreated, id.String()),
		ProxyID:   id.String(),
		Name:      name,
		Host:      host,
		Port:      port,
		Protocol:  protocol.String(),
	})

	return proxy, nil
}

// SetCredentials sets authentication credentials for the proxy
func (p *Proxy) SetCredentials(credentials ProxyCredentials) {
	p.Credentials = &credentials
	p.UpdatedAt = shared.NewTimestamp(time.Now())
}

// RemoveCredentials removes authentication credentials
func (p *Proxy) RemoveCredentials() {
	p.Credentials = nil
	p.UpdatedAt = shared.NewTimestamp(time.Now())
}

// UpdateStatus updates the proxy status
func (p *Proxy) UpdateStatus(status ProxyStatus) {
	oldStatus := p.Status
	p.Status = status
	p.UpdatedAt = shared.NewTimestamp(time.Now())

	if oldStatus.String() != status.String() {
		p.addEvent(ProxyStatusChanged{
			BaseEvent: shared.NewBaseEvent(EventProxyStatusChanged, p.Id.String()),
			ProxyID:   p.Id.String(),
			OldStatus: oldStatus.String(),
			NewStatus: status.String(),
		})
	}
}

// MarkChecked updates the last checked timestamp
func (p *Proxy) MarkChecked() {
	now := shared.NewTimestamp(time.Now())
	p.LastCheckedAt = &now
	p.UpdatedAt = shared.NewTimestamp(time.Now())
}

// RecordSuccess records a successful proxy connection
func (p *Proxy) RecordSuccess(latencyMs int) {
	p.SuccessCount++
	p.AverageLatency = ((p.AverageLatency * (p.SuccessCount - 1)) + latencyMs) / p.SuccessCount
	p.UpdateStatus(ProxyStatusActive)
	p.MarkChecked()
}

// RecordFailure records a failed proxy connection
func (p *Proxy) RecordFailure(reason string) {
	p.FailureCount++
	now := shared.NewTimestamp(time.Now())
	p.LastFailureAt = &now
	p.UpdateStatus(ProxyStatusFailed)
	p.MarkChecked()

	p.addEvent(ProxyFailed{
		BaseEvent: shared.NewBaseEvent(EventProxyFailed, p.Id.String()),
		ProxyID:   p.Id.String(),
		Reason:    reason,
		FailedAt:  now.Time(),
	})
}

// Activate activates the proxy
func (p *Proxy) Activate() {
	p.UpdateStatus(ProxyStatusActive)
}

// Deactivate deactivates the proxy
func (p *Proxy) Deactivate() {
	p.UpdateStatus(ProxyStatusInactive)
}

// AddTag adds a tag to the proxy
func (p *Proxy) AddTag(tagID tag.TagID) error {
	// Check if tag already exists
	for _, existingTag := range p.Tags {
		if existingTag.String() == tagID.String() {
			return errors.New("tag already exists on proxy")
		}
	}
	p.Tags = append(p.Tags, tagID)
	p.UpdatedAt = shared.NewTimestamp(time.Now())
	return nil
}

// RemoveTag removes a tag from the proxy
func (p *Proxy) RemoveTag(tagID tag.TagID) {
	for i, existingTag := range p.Tags {
		if existingTag.String() == tagID.String() {
			p.Tags = append(p.Tags[:i], p.Tags[i+1:]...)
			p.UpdatedAt = shared.NewTimestamp(time.Now())
			return
		}
	}
}

// HasTag checks if the proxy has a specific tag
func (p *Proxy) HasTag(tagID tag.TagID) bool {
	for _, existingTag := range p.Tags {
		if existingTag.String() == tagID.String() {
			return true
		}
	}
	return false
}

// GetConnectionURL returns the full proxy connection URL
func (p *Proxy) GetConnectionURL() string {
	if p.Credentials != nil {
		return fmt.Sprintf("%s://%s:%s@%s:%d",
			p.Protocol.String(),
			p.Credentials.Username,
			p.Credentials.Password,
			p.Host,
			p.Port)
	}
	return fmt.Sprintf("%s://%s:%d", p.Protocol.String(), p.Host, p.Port)
}

func (p *Proxy) addEvent(event shared.DomainEvent) {
	p.Events = append(p.Events, event)
}

func (p *Proxy) ClearEvents() {
	p.Events = make([]shared.DomainEvent, 0)
}
