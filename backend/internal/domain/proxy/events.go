package proxy

import (
	"parrotflow/internal/domain/shared"
	"time"
)

const (
	EventProxyCreated       = "proxy.created"
	EventProxyStatusChanged = "proxy.status.changed"
	EventProxyFailed        = "proxy.failed"
)

type ProxyCreated struct {
	shared.BaseEvent
	ProxyID  string
	Name     string
	Host     string
	Port     int
	Protocol string
}

type ProxyStatusChanged struct {
	shared.BaseEvent
	ProxyID   string
	OldStatus string
	NewStatus string
}

type ProxyFailed struct {
	shared.BaseEvent
	ProxyID  string
	Reason   string
	FailedAt time.Time
}
