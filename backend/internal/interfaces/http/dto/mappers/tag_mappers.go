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

// Mapper functions using functional approach instead of empty structs

func TagToCreateResponse(t *tag.Tag) *commands.CreateTagResponse {
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

func TagToUpdateResponse(t *tag.Tag) *commands.UpdateTagResponse {
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

func TagToDeleteResponse() *commands.DeleteTagResponse {
	response := &commands.DeleteTagResponse{}
	response.Body.Success = true
	return response
}

func TagToGetResponse(t *tag.Tag) *queries.GetTagResponse {
	response := &queries.GetTagResponse{}
	response.Body = buildTagDTO(t)
	return response
}

func TagToListResponse(tags []*tag.Tag) *queries.ListTagsResponse {
	response := &queries.ListTagsResponse{}
	response.Body.Tags = MapSlicePtr(tags, buildTagDTO)
	return response
}

// Mapper instances for handler injection - using functional types
var (
	TagCreateMapper = CreateMapperFunc[*tag.Tag, *commands.CreateTagResponse](TagToCreateResponse)
	TagUpdateMapper = UpdateMapperFunc[*tag.Tag, *commands.UpdateTagResponse](TagToUpdateResponse)
	TagDeleteMapper = DeleteMapperFunc[*commands.DeleteTagResponse](TagToDeleteResponse)
	TagGetMapper    = GetMapperFunc[*tag.Tag, *queries.GetTagResponse](TagToGetResponse)
	TagListMapper   = ListMapperFunc[tag.Tag, *queries.ListTagsResponse](TagToListResponse)
)
