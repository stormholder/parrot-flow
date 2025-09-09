package scenario

import "parrotflow/internal/api"

type ScenarioQuery struct {
	Name string   `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	api.PageQuery
	api.OrderByQuery
}
