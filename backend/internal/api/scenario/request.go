package scenario

import "parrotflow/pkg/shared"

type ScenarioQuery struct {
	Name string   `json:"name,omitempty"`
	Tags []string `json:"tags,omitempty"`
	shared.PageQuery
	shared.OrderByQuery
}
