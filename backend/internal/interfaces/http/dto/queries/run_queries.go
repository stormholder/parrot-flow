package queries

type GetRunRequest struct {
	ID string `path:"id"`
}

type GetRunResponse struct {
	Body struct {
		ID         string  `json:"id"`
		ScenarioID string  `json:"scenario_id"`
		Status     string  `json:"status"`
		Parameters string  `json:"parameters"`
		StartedAt  *string `json:"started_at,omitempty"`
		FinishedAt *string `json:"finished_at,omitempty"`
		CreatedAt  string  `json:"created_at"`
		UpdatedAt  string  `json:"updated_at"`
	}
}

type ListRunsRequest struct {
	ScenarioID string `query:"scenario_id"`
	Status     string `query:"status"`
	Page       int    `query:"page"`
	RPP        int    `query:"rpp"`
}

type RunListItem struct {
	ID         string  `json:"id"`
	ScenarioID string  `json:"scenario_id"`
	Status     string  `json:"status"`
	Parameters string  `json:"parameters"`
	StartedAt  *string `json:"started_at,omitempty"`
	FinishedAt *string `json:"finished_at,omitempty"`
	CreatedAt  string  `json:"created_at"`
}

type ListRunsResponse struct {
	Body struct {
		Data  []RunListItem `json:"data"`
		Total int           `json:"total"`
		Page  int           `json:"page"`
		RPP   int           `json:"rpp"`
	}
}
