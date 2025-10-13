package models

import "time"

// Agent represents a browser automation agent in the database
type Agent struct {
	Model
	Name             string     `json:"name" gorm:"size:255;not null;uniqueIndex"`
	Status           string     `json:"status" gorm:"size:20;not null;index"`
	Capabilities     string     `json:"capabilities" gorm:"type:jsonb;not null"` // JSON
	CurrentRunCount  int        `json:"current_run_count" gorm:"default:0"`
	LastHeartbeatAt  *time.Time `json:"last_heartbeat_at,omitempty" gorm:"index"`
	RegisteredAt     time.Time  `json:"registered_at" gorm:"not null"`
	ConnectionInfo   string     `json:"connection_info" gorm:"type:jsonb;not null"` // JSON
	Metadata         string     `json:"metadata,omitempty" gorm:"type:jsonb"`        // JSON
	Tags             []Tag      `json:"tags" gorm:"many2many:agent_tags;"`
}

// TableName specifies the table name for GORM
func (Agent) TableName() string {
	return "agents"
}
