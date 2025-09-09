package models

type ScenarioBase struct {
	Model
	Name        string `json:"name" gorm:"size: 255;not null"`
	Description string `json:"description,omitempty" gorm:"default:NULL"`
	Tag         string `json:"tag,omitempty" gorm:"default:NULL"`
	Icon        string `json:"icon,omitempty" gorm:"default:NULL"`
}

type Scenario struct {
	ScenarioBase
	Payload    string `json:"payload" gorm:"not null"`
	InputData  string `json:"input_data" gorm:"not null"`
	Parameters string `json:"parameters" gorm:"not null"`
}
