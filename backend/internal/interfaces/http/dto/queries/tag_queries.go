package queries

type GetTagRequest struct {
	ID string `path:"id"`
}

type GetTagResponse struct {
	Body struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Category    string `json:"category"`
		Description string `json:"description"`
		Color       string `json:"color"`
		IsSystem    bool   `json:"is_system"`
		CreatedAt   string `json:"created_at"`
		UpdatedAt   string `json:"updated_at"`
	}
}

type ListTagsRequest struct {
	Category string `query:"category" enum:"system,custom,capacity,region" doc:"Filter by category (optional)"`
}

type ListTagsResponse struct {
	Body struct {
		Tags []TagDTO `json:"tags"`
	}
}

type TagDTO struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Color       string `json:"color"`
	IsSystem    bool   `json:"is_system"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}
