package commands

type CreateRunRequest struct {
	Body struct {
		ScenarioID string `json:"scenario_id"`
		Parameters string `json:"parameters"`
	}
}

type CreateRunResponse struct {
	Body struct {
		ID         string `json:"id"`
		ScenarioID string `json:"scenario_id"`
		Status     string `json:"status"`
		Parameters string `json:"parameters"`
		CreatedAt  string `json:"created_at"`
	}
}

type StartRunRequest struct {
	ID string `path:"id"`
}

type StartRunResponse struct {
	Body struct {
		ID        string `json:"id"`
		Status    string `json:"status"`
		StartedAt string `json:"started_at"`
	}
}
