package mappers

import (
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/queries"
)

func ToCreateTagResponse(t *tag.Tag) *commands.CreateTagResponse {
	response := &commands.CreateTagResponse{}
	response.Body.ID = t.Id.String()
	response.Body.Name = t.Name
	response.Body.Category = t.Category.String()
	response.Body.Description = t.Description
	response.Body.Color = t.Color
	response.Body.IsSystem = t.IsSystem
	response.Body.CreatedAt = t.CreatedAt.Time().Format("2006-01-02T15:04:05Z07:00")
	response.Body.UpdatedAt = t.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00")
	return response
}

func ToUpdateTagResponse(t *tag.Tag) *commands.UpdateTagResponse {
	response := &commands.UpdateTagResponse{}
	response.Body.ID = t.Id.String()
	response.Body.Name = t.Name
	response.Body.Category = t.Category.String()
	response.Body.Description = t.Description
	response.Body.Color = t.Color
	response.Body.IsSystem = t.IsSystem
	response.Body.UpdatedAt = t.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00")
	return response
}

func ToGetTagResponse(t *tag.Tag) *queries.GetTagResponse {
	response := &queries.GetTagResponse{}
	response.Body.ID = t.Id.String()
	response.Body.Name = t.Name
	response.Body.Category = t.Category.String()
	response.Body.Description = t.Description
	response.Body.Color = t.Color
	response.Body.IsSystem = t.IsSystem
	response.Body.CreatedAt = t.CreatedAt.Time().Format("2006-01-02T15:04:05Z07:00")
	response.Body.UpdatedAt = t.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00")
	return response
}

func ToListTagsResponse(tags []*tag.Tag) *queries.ListTagsResponse {
	response := &queries.ListTagsResponse{}
	response.Body.Tags = make([]queries.TagDTO, len(tags))
	for i, t := range tags {
		response.Body.Tags[i] = queries.TagDTO{
			ID:          t.Id.String(),
			Name:        t.Name,
			Category:    t.Category.String(),
			Description: t.Description,
			Color:       t.Color,
			IsSystem:    t.IsSystem,
			CreatedAt:   t.CreatedAt.Time().Format("2006-01-02T15:04:05Z07:00"),
			UpdatedAt:   t.UpdatedAt.Time().Format("2006-01-02T15:04:05Z07:00"),
		}
	}
	return response
}
