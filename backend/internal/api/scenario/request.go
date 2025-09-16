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

type ScenarioPatch struct {
	Name        *string                 `json:"name,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Tag         *string                 `json:"tag,omitempty"`
	Icon        *string                 `json:"icon,omitempty"`
	Context     *ScenarioContext        `json:"context,omitempty"`
	InputData   *interface{}            `json:"input_data,omitempty"`
	Parameters  *ScenarioNodeParameters `json:"parameters,omitempty"`
}

type ScenarioPatchRequest struct {
	GetScenarioByIDRequest
	Body struct {
		ScenarioPatch
	}
}
