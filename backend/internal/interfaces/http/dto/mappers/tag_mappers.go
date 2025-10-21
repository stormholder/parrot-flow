package mappers

import (
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/queries"
)

func buildTagDTO(t *tag.Tag) queries.TagDTO {
	return queries.TagDTO{
		ID:          t.Id.String(),
		Name:        t.Name,
		Category:    t.Category.String(),
		Description: t.Description,
		Color:       t.Color,
		IsSystem:    t.IsSystem,
		CreatedAt:   FormatTimestamp(t.CreatedAt.Time()),
		UpdatedAt:   FormatTimestamp(t.UpdatedAt.Time()),
	}
}

type TagCreateMapper struct{}

func (m TagCreateMapper) Map(t *tag.Tag) *commands.CreateTagResponse {
	dto := buildTagDTO(t)
	response := &commands.CreateTagResponse{}
	response.Body.ID = dto.ID
	response.Body.Name = dto.Name
	response.Body.Category = dto.Category
	response.Body.Description = dto.Description
	response.Body.Color = dto.Color
	response.Body.IsSystem = dto.IsSystem
	response.Body.CreatedAt = dto.CreatedAt
	response.Body.UpdatedAt = dto.UpdatedAt
	return response
}

// TagUpdateMapper maps Tag to UpdateTagResponse
type TagUpdateMapper struct{}

func (m TagUpdateMapper) Map(t *tag.Tag) *commands.UpdateTagResponse {
	dto := buildTagDTO(t)
	response := &commands.UpdateTagResponse{}
	response.Body.ID = dto.ID
	response.Body.Name = dto.Name
	response.Body.Category = dto.Category
	response.Body.Description = dto.Description
	response.Body.Color = dto.Color
	response.Body.IsSystem = dto.IsSystem
	response.Body.UpdatedAt = dto.UpdatedAt
	return response
}

type TagDeleteMapper struct{}

func (m TagDeleteMapper) Map() *commands.DeleteTagResponse {
	response := &commands.DeleteTagResponse{}
	response.Body.Success = true
	return response
}

type TagGetMapper struct{}

func (m TagGetMapper) Map(t *tag.Tag) *queries.GetTagResponse {
	response := &queries.GetTagResponse{}
	response.Body = buildTagDTO(t)
	return response
}

type TagListMapper struct{}

func (m TagListMapper) Map(tags []*tag.Tag) *queries.ListTagsResponse {
	response := &queries.ListTagsResponse{}
	response.Body.Tags = MapSlicePtr(tags, buildTagDTO)
	return response
}

func ToCreateTagResponse(t *tag.Tag) *commands.CreateTagResponse {
	return TagCreateMapper{}.Map(t)
}

func ToUpdateTagResponse(t *tag.Tag) *commands.UpdateTagResponse {
	return TagUpdateMapper{}.Map(t)
}

func ToGetTagResponse(t *tag.Tag) *queries.GetTagResponse {
	return TagGetMapper{}.Map(t)
}

func ToListTagsResponse(tags []*tag.Tag) *queries.ListTagsResponse {
	return TagListMapper{}.Map(tags)
}
