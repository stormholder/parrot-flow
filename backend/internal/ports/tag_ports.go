package ports

import (
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/models"
)

func TagParseID(id string) uint64 {
	return parseID(id)
}

func TagDomainEntityToPersistence(t *tag.Tag) (*models.Tag, error) {
	model := &models.Tag{
		Model: models.Model{
			ID:        parseID(t.Id.String()),
			CreatedAt: t.CreatedAt.Time(),
			UpdatedAt: t.UpdatedAt.Time(),
		},
		Name:        t.Name,
		Category:    t.Category.String(),
		Description: t.Description,
		Color:       t.Color,
		IsSystem:    t.IsSystem,
	}
	return model, nil
}

func TagPersistenceToDomainEntity(model *models.Tag) (*tag.Tag, error) {
	tagID, err := tag.NewTagID(formatID(model.ID))
	if err != nil {
		return nil, err
	}

	category, err := tag.NewTagCategory(model.Category)
	if err != nil {
		return nil, err
	}

	var t *tag.Tag
	if model.IsSystem {
		t, err = tag.NewSystemTag(tagID, model.Name, category, model.Description)
	} else {
		t, err = tag.NewTag(tagID, model.Name, category)
		if err != nil {
			return nil, err
		}
		t.UpdateDescription(model.Description)
	}
	if err != nil {
		return nil, err
	}

	if model.Color != "" {
		if err := t.UpdateColor(model.Color); err != nil {
			return nil, err
		}
	}

	return t, nil
}
