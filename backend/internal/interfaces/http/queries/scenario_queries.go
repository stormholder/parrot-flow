package queries

type GetScenarioRequest struct {
	ID string `path:"id"`
}

type ScenarioResponseItem struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Tag         string `json:"tag"`
	Icon        string `json:"icon"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type GetScenarioResponse struct {
	Body ScenarioResponseItem
}

type ListScenariosRequest struct {
	Name string `query:"name"`
	Tag  string `query:"tag"`
	Page int    `query:"page"`
	RPP  int    `query:"rpp"`
}

type ListScenariosResponse struct {
	Body struct {
		Data  []ScenarioResponseItem `json:"data"`
		Total int                    `json:"total"`
		Page  int                    `json:"page"`
		RPP   int                    `json:"rpp"`
	}
}
