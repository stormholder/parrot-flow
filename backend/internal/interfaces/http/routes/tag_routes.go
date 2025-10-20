package routes

import (
	"github.com/danielgtaylor/huma/v2"
	"parrotflow/internal/interfaces/http/handlers"
)

func RegisterTagRoutes(api *huma.API, tagHandler *handlers.TagHandler) {

	huma.Register(*api, huma.Operation{
		OperationID: "create-tag",
		Method:      "POST",
		Path:        "/api/tags/",
		Summary:     "Create a new tag",
		Description: "Create a new tag for labeling entities",
		Tags:        []string{"tags"},
	}, tagHandler.CreateTag)

	huma.Register(*api, huma.Operation{
		OperationID: "get-tag",
		Method:      "GET",
		Path:        "/api/tags/{id}",
		Summary:     "Get a tag by ID",
		Description: "Retrieve a specific tag by its ID",
		Tags:        []string{"tags"},
	}, tagHandler.GetTag)

	huma.Register(*api, huma.Operation{
		OperationID: "list-tags",
		Method:      "GET",
		Path:        "/api/tags/",
		Summary:     "List tags",
		Description: "Get a list of tags with optional category filtering",
		Tags:        []string{"tags"},
	}, tagHandler.ListTags)

	huma.Register(*api, huma.Operation{
		OperationID: "update-tag",
		Method:      "PATCH",
		Path:        "/api/tags/{id}",
		Summary:     "Update a tag",
		Description: "Update an existing tag (description and color only)",
		Tags:        []string{"tags"},
	}, tagHandler.UpdateTag)

	huma.Register(*api, huma.Operation{
		OperationID: "delete-tag",
		Method:      "DELETE",
		Path:        "/api/tags/{id}",
		Summary:     "Delete a tag",
		Description: "Delete a tag by its ID (system tags cannot be deleted)",
		Tags:        []string{"tags"},
	}, tagHandler.DeleteTag)
}
