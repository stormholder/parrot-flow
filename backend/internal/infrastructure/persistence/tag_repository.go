package persistence

import (
	"context"
	"errors"
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/models"
	"parrotflow/internal/ports"

	"gorm.io/gorm"
)

type TagRepository struct {
	db *gorm.DB
}

func NewTagRepository(db *gorm.DB) *TagRepository {
	return &TagRepository{db: db}
}

func (r *TagRepository) Save(ctx context.Context, t *tag.Tag) error {
	model, err := ports.TagDomainEntityToPersistence(t)
	if err != nil {
		return err
	}

	if err = r.db.WithContext(ctx).Save(model).Error; err != nil {
		return err
	}

	return nil
}

func (r *TagRepository) FindByID(ctx context.Context, id tag.TagID) (*tag.Tag, error) {
	var model models.Tag
	if err := r.db.WithContext(ctx).Where("id = ?", ports.TagParseID(id.String())).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("tag not found")
		}
		return nil, err
	}

	return ports.TagPersistenceToDomainEntity(&model)
}

func (r *TagRepository) FindByName(ctx context.Context, name string) (*tag.Tag, error) {
	var model models.Tag
	if err := r.db.WithContext(ctx).Where("LOWER(name) = LOWER(?)", name).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("tag not found")
		}
		return nil, err
	}

	return ports.TagPersistenceToDomainEntity(&model)
}

func (r *TagRepository) FindByCategory(ctx context.Context, category tag.TagCategory) ([]*tag.Tag, error) {
	var models []models.Tag
	if err := r.db.WithContext(ctx).Where("category = ?", category.String()).Find(&models).Error; err != nil {
		return nil, err
	}

	tags := make([]*tag.Tag, len(models))
	for i, model := range models {
		t, err := ports.TagPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		tags[i] = t
	}

	return tags, nil
}

func (r *TagRepository) FindAll(ctx context.Context) ([]*tag.Tag, error) {
	var models []models.Tag
	if err := r.db.WithContext(ctx).Order("category ASC, name ASC").Find(&models).Error; err != nil {
		return nil, err
	}

	tags := make([]*tag.Tag, len(models))
	for i, model := range models {
		t, err := ports.TagPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		tags[i] = t
	}

	return tags, nil
}

func (r *TagRepository) FindByIDs(ctx context.Context, ids []tag.TagID) ([]*tag.Tag, error) {
	if len(ids) == 0 {
		return []*tag.Tag{}, nil
	}

	parsedIDs := make([]uint64, len(ids))
	for i, id := range ids {
		parsedIDs[i] = ports.TagParseID(id.String())
	}

	var models []models.Tag
	if err := r.db.WithContext(ctx).Where("id IN ?", parsedIDs).Find(&models).Error; err != nil {
		return nil, err
	}

	tags := make([]*tag.Tag, len(models))
	for i, model := range models {
		t, err := ports.TagPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		tags[i] = t
	}

	return tags, nil
}

func (r *TagRepository) Delete(ctx context.Context, id tag.TagID) error {
	return r.db.WithContext(ctx).Where("id = ?", ports.TagParseID(id.String())).Delete(&models.Tag{}).Error
}

func (r *TagRepository) Exists(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Tag{}).Where("LOWER(name) = LOWER(?)", name).Count(&count).Error
	return count > 0, err
}
