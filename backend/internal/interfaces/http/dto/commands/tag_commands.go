package commands

type CreateTagRequest struct {
	Body struct {
		Name        string `json:"name" minLength:"1" maxLength:"100" doc:"Tag name (lowercase, unique)"`
		Category    string `json:"category" enum:"system,custom,capacity,region" doc:"Tag category"`
		Description string `json:"description,omitempty" doc:"Tag description"`
		Color       string `json:"color,omitempty" pattern:"^#[0-9A-Fa-f]{6}$" doc:"Hex color (#RRGGBB)"`
	}
}

type CreateTagResponse struct {
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

type UpdateTagRequest struct {
	ID   string `path:"id"`
	Body struct {
		Description *string `json:"description,omitempty" doc:"Tag description"`
		Color       *string `json:"color,omitempty" pattern:"^#[0-9A-Fa-f]{6}$" doc:"Hex color (#RRGGBB)"`
	}
}

type UpdateTagResponse struct {
	Body struct {
		ID          string `json:"id"`
		Name        string `json:"name"`
		Category    string `json:"category"`
		Description string `json:"description"`
		Color       string `json:"color"`
		IsSystem    bool   `json:"is_system"`
		UpdatedAt   string `json:"updated_at"`
	}
}

type DeleteTagRequest struct {
	ID string `path:"id"`
}

type DeleteTagResponse struct {
	Body struct {
		Success bool `json:"success"`
	}
}
