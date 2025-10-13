package models

import "time"

// Proxy represents a proxy server in the database
type Proxy struct {
	Model
	Name            string     `json:"name" gorm:"size:255;not null;uniqueIndex"`
	Host            string     `json:"host" gorm:"size:255;not null"`
	Port            int        `json:"port" gorm:"not null"`
	Protocol        string     `json:"protocol" gorm:"size:10;not null"` // http, https, socks5
	Username        string     `json:"username,omitempty" gorm:"size:255"`
	Password        string     `json:"password,omitempty" gorm:"size:255"` // Should be encrypted
	Status          string     `json:"status" gorm:"size:20;not null;index"`
	LastCheckedAt   *time.Time `json:"last_checked_at,omitempty"`
	LastFailureAt   *time.Time `json:"last_failure_at,omitempty"`
	FailureCount    int        `json:"failure_count" gorm:"default:0"`
	SuccessCount    int        `json:"success_count" gorm:"default:0"`
	AverageLatency  int        `json:"average_latency" gorm:"default:0"` // milliseconds
	Tags            []Tag      `json:"tags" gorm:"many2many:proxy_tags;"`
}

// TableName specifies the table name for GORM
func (Proxy) TableName() string {
	return "proxies"
}
