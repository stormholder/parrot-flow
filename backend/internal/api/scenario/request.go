package scenario

import "parrotflow/internal/api"

type ScenarioQuery struct {
	Name string   `json:"name,omitempty" query:"name"`
	Tags []string `json:"tags,omitempty" query:"tags"`
	api.PageQuery
	api.OrderByQuery
}

type GetScenarioByIDRequest struct {
	ID uint `path:"id" required:"true"`
}
