package commands

type CreateScenarioRequest struct {
	Body struct {
		Name        string `json:"name"`
		Description string `json:"description,omitempty"`
		Tag         string `json:"tag,omitempty"`
		Icon        string `json:"icon,omitempty"`
	}
}

type CreateScenarioResponse struct {
	Body struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Tag         string `json:"tag"`
		Icon        string `json:"icon"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
}

type UpdateScenarioRequest struct {
	ID   string `path:"id"`
	Body struct {
		Name        *string `json:"name,omitempty"`
		Description *string `json:"description,omitempty"`
		Tag         *string `json:"tag,omitempty"`
		Icon        *string `json:"icon,omitempty"`
	}
}

type UpdateScenarioResponse struct {
	Body struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Description string `json:"description"`
		Tag         string `json:"tag"`
		Icon        string `json:"icon"`
		UpdatedAt   string `json:"updated_at"`
	}
}

type DeleteScenarioRequest struct {
	ID string `path:"id"`
}

type DeleteScenarioResponse struct {
	Body struct {
		Success bool `json:"success"`
	}
}
