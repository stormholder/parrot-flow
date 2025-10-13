package handlers

import (
	"context"
	command "parrotflow/internal/application/command/tag"
	query "parrotflow/internal/application/query/tag"
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/interfaces/http/dto/commands"
	"parrotflow/internal/interfaces/http/dto/mappers"
	"parrotflow/internal/interfaces/http/dto/queries"
)

type TagHandler struct {
	createCommandHandler *command.CreateTagCommandHandler
	updateCommandHandler *command.UpdateTagCommandHandler
	deleteCommandHandler *command.DeleteTagCommandHandler
	getQueryHandler      *query.GetTagQueryHandler
	listQueryHandler     *query.ListTagsQueryHandler
}

func NewTagHandler(
	createCommandHandler *command.CreateTagCommandHandler,
	updateCommandHandler *command.UpdateTagCommandHandler,
	deleteCommandHandler *command.DeleteTagCommandHandler,
	getQueryHandler *query.GetTagQueryHandler,
	listQueryHandler *query.ListTagsQueryHandler,
) *TagHandler {
	return &TagHandler{
		createCommandHandler: createCommandHandler,
		updateCommandHandler: updateCommandHandler,
		deleteCommandHandler: deleteCommandHandler,
		getQueryHandler:      getQueryHandler,
		listQueryHandler:     listQueryHandler,
	}
}

func (h *TagHandler) CreateTag(ctx context.Context, req *commands.CreateTagRequest) (*commands.CreateTagResponse, error) {
	cmd := command.CreateTagCommand{
		Name:        req.Body.Name,
		Category:    req.Body.Category,
		Description: req.Body.Description,
		Color:       req.Body.Color,
	}

	t, err := h.createCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return mappers.ToCreateTagResponse(t), nil
}

func (h *TagHandler) GetTag(ctx context.Context, req *queries.GetTagRequest) (*queries.GetTagResponse, error) {
	tagID, err := tag.NewTagID(req.ID)
	if err != nil {
		return nil, err
	}

	q := query.GetTagQuery{
		ID: tagID,
	}

	t, err := h.getQueryHandler.Handle(ctx, q)
	if err != nil {
		return nil, err
	}

	return mappers.ToGetTagResponse(t), nil
}

func (h *TagHandler) ListTags(ctx context.Context, req *queries.ListTagsRequest) (*queries.ListTagsResponse, error) {
	q := query.ListTagsQuery{}

	if req.Category != "" {
		category, err := tag.NewTagCategory(req.Category)
		if err != nil {
			return nil, err
		}
		q.Category = &category
	}

	tags, err := h.listQueryHandler.Handle(ctx, q)
	if err != nil {
		return nil, err
	}

	return mappers.ToListTagsResponse(tags), nil
}

func (h *TagHandler) UpdateTag(ctx context.Context, req *commands.UpdateTagRequest) (*commands.UpdateTagResponse, error) {
	tagID, err := tag.NewTagID(req.ID)
	if err != nil {
		return nil, err
	}

	cmd := command.UpdateTagCommand{
		ID:          tagID,
		Description: req.Body.Description,
		Color:       req.Body.Color,
	}

	t, err := h.updateCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	return mappers.ToUpdateTagResponse(t), nil
}

func (h *TagHandler) DeleteTag(ctx context.Context, req *commands.DeleteTagRequest) (*commands.DeleteTagResponse, error) {
	tagID, err := tag.NewTagID(req.ID)
	if err != nil {
		return nil, err
	}

	cmd := command.DeleteTagCommand{
		ID: tagID,
	}

	err = h.deleteCommandHandler.Handle(ctx, cmd)
	if err != nil {
		return nil, err
	}

	response := &commands.DeleteTagResponse{}
	response.Body.Success = true

	return response, nil
}
