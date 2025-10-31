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
	// Command handlers
	createCommandHandler *command.CreateTagCommandHandler
	updateCommandHandler *command.UpdateTagCommandHandler
	deleteCommandHandler *command.DeleteTagCommandHandler

	// Query handlers
	getQueryHandler  *query.GetTagQueryHandler
	listQueryHandler *query.ListTagsQueryHandler

	// Mappers - using functional types
	createMapper mappers.CreateMapperFunc[*tag.Tag, *commands.CreateTagResponse]
	updateMapper mappers.UpdateMapperFunc[*tag.Tag, *commands.UpdateTagResponse]
	deleteMapper mappers.DeleteMapperFunc[*commands.DeleteTagResponse]
	getMapper    mappers.GetMapperFunc[*tag.Tag, *queries.GetTagResponse]
	listMapper   mappers.ListMapperFunc[tag.Tag, *queries.ListTagsResponse]
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
		createMapper:         mappers.TagCreateMapper,
		updateMapper:         mappers.TagUpdateMapper,
		deleteMapper:         mappers.TagDeleteMapper,
		getMapper:            mappers.TagGetMapper,
		listMapper:           mappers.TagListMapper,
	}
}

func (h *TagHandler) CreateTag(ctx context.Context, req *commands.CreateTagRequest) (*commands.CreateTagResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.CreateTagRequest) (command.CreateTagCommand, error) {
			return command.CreateTagCommand{
				Name:        r.Body.Name,
				Category:    r.Body.Category,
				Description: r.Body.Description,
				Color:       r.Body.Color,
			}, nil
		},
		CommandHandlerFunc[command.CreateTagCommand, *tag.Tag](h.createCommandHandler.Handle),
		h.createMapper,
	)
}

func (h *TagHandler) UpdateTag(ctx context.Context, req *commands.UpdateTagRequest) (*commands.UpdateTagResponse, error) {
	return HandleCommand(
		ctx,
		req,
		func(r *commands.UpdateTagRequest) (command.UpdateTagCommand, error) {
			tagID, err := tag.NewTagID(r.ID)
			if err != nil {
				return command.UpdateTagCommand{}, err
			}
			return command.UpdateTagCommand{
				ID:          tagID,
				Description: r.Body.Description,
				Color:       r.Body.Color,
			}, nil
		},
		CommandHandlerFunc[command.UpdateTagCommand, *tag.Tag](h.updateCommandHandler.Handle),
		h.updateMapper,
	)
}

func (h *TagHandler) DeleteTag(ctx context.Context, req *commands.DeleteTagRequest) (*commands.DeleteTagResponse, error) {
	return HandleSimpleCommand(
		ctx,
		req,
		func(r *commands.DeleteTagRequest) (command.DeleteTagCommand, error) {
			tagID, err := tag.NewTagID(r.ID)
			if err != nil {
				return command.DeleteTagCommand{}, err
			}
			return command.DeleteTagCommand{ID: tagID}, nil
		},
		SimpleCommandHandlerFunc[command.DeleteTagCommand](h.deleteCommandHandler.Handle),
		h.deleteMapper.Map,
	)
}

func (h *TagHandler) GetTag(ctx context.Context, req *queries.GetTagRequest) (*queries.GetTagResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.GetTagRequest) (query.GetTagQuery, error) {
			tagID, err := tag.NewTagID(r.ID)
			if err != nil {
				return query.GetTagQuery{}, err
			}
			return query.GetTagQuery{ID: tagID}, nil
		},
		QueryHandlerFunc[query.GetTagQuery, *tag.Tag](h.getQueryHandler.Handle),
		h.getMapper,
	)
}

func (h *TagHandler) ListTags(ctx context.Context, req *queries.ListTagsRequest) (*queries.ListTagsResponse, error) {
	return HandleQuery(
		ctx,
		req,
		func(r *queries.ListTagsRequest) (query.ListTagsQuery, error) {
			q := query.ListTagsQuery{}
			if r.Category != "" {
				category, err := tag.NewTagCategory(r.Category)
				if err != nil {
					return query.ListTagsQuery{}, err
				}
				q.Category = &category
			}
			return q, nil
		},
		QueryHandlerFunc[query.ListTagsQuery, []*tag.Tag](h.listQueryHandler.Handle),
		h.listMapper,
	)
}
