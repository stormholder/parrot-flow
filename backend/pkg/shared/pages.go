package shared

type Pages struct {
	Total       int         `json:"total"`
	Pages       int         `json:"pages"`
	PerPage     int         `json:"per_page"`
	CurrentPage int         `json:"current_page"`
	Data        interface{} `json:"data"`
}

type PageQuery struct {
	Page    int `json:"page,omitempty"`
	PerPage int `json:"per_page,omitempty"`
}

type OrderByQuery struct {
	Field     string `json:"field,omitempty"`
	Direction string `json:"direction,omitempty"`
}
