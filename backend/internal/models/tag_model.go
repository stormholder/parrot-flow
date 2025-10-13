package models

// Tag represents a tag in the database
type Tag struct {
	Model
	Name        string `json:"name" gorm:"size:100;not null;uniqueIndex"`
	Category    string `json:"category" gorm:"size:50;not null;index"`
	Description string `json:"description,omitempty" gorm:"type:text"`
	Color       string `json:"color,omitempty" gorm:"size:7"` // #RRGGBB
	IsSystem    bool   `json:"is_system" gorm:"not null;default:false;index"`
}

// TableName specifies the table name for GORM
func (Tag) TableName() string {
	return "tags"
}
