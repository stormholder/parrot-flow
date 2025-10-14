package persistence

import (
	"context"
	"errors"
	"parrotflow/internal/domain/proxy"
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/models"
	"parrotflow/internal/ports"

	"gorm.io/gorm"
)

type ProxyRepository struct {
	db *gorm.DB
}

func NewProxyRepository(db *gorm.DB) *ProxyRepository {
	return &ProxyRepository{db: db}
}

func (r *ProxyRepository) Save(ctx context.Context, p *proxy.Proxy) error {
	model, err := ports.ProxyDomainEntityToPersistence(p)
	if err != nil {
		return err
	}

	// Load tags from IDs
	if len(p.Tags) > 0 {
		tagIDs := make([]uint64, len(p.Tags))
		for i, tagID := range p.Tags {
			tagIDs[i] = ports.TagParseID(tagID.String())
		}
		var tags []models.Tag
		if err := r.db.WithContext(ctx).Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
			return err
		}
		model.Tags = tags
	}

	// Use Association for many-to-many relationship
	if err = r.db.WithContext(ctx).Save(model).Error; err != nil {
		return err
	}

	// Update associations
	if err = r.db.WithContext(ctx).Model(model).Association("Tags").Replace(model.Tags); err != nil {
		return err
	}

	return nil
}

func (r *ProxyRepository) FindByID(ctx context.Context, id proxy.ProxyID) (*proxy.Proxy, error) {
	var model models.Proxy
	if err := r.db.WithContext(ctx).Preload("Tags").Where("id = ?", ports.ProxyParseID(id.String())).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("proxy not found")
		}
		return nil, err
	}

	return ports.ProxyPersistenceToDomainEntity(&model)
}

func (r *ProxyRepository) FindByName(ctx context.Context, name string) (*proxy.Proxy, error) {
	var model models.Proxy
	if err := r.db.WithContext(ctx).Preload("Tags").Where("name = ?", name).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("proxy not found")
		}
		return nil, err
	}

	return ports.ProxyPersistenceToDomainEntity(&model)
}

func (r *ProxyRepository) FindAll(ctx context.Context) ([]*proxy.Proxy, error) {
	var models []models.Proxy
	if err := r.db.WithContext(ctx).Preload("Tags").Order("name ASC").Find(&models).Error; err != nil {
		return nil, err
	}

	proxies := make([]*proxy.Proxy, len(models))
	for i, model := range models {
		p, err := ports.ProxyPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		proxies[i] = p
	}

	return proxies, nil
}

func (r *ProxyRepository) FindByStatus(ctx context.Context, status proxy.ProxyStatus) ([]*proxy.Proxy, error) {
	var models []models.Proxy
	if err := r.db.WithContext(ctx).Preload("Tags").Where("status = ?", status.String()).Find(&models).Error; err != nil {
		return nil, err
	}

	proxies := make([]*proxy.Proxy, len(models))
	for i, model := range models {
		p, err := ports.ProxyPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		proxies[i] = p
	}

	return proxies, nil
}

func (r *ProxyRepository) FindByTags(ctx context.Context, tagIDs []tag.TagID) ([]*proxy.Proxy, error) {
	if len(tagIDs) == 0 {
		return []*proxy.Proxy{}, nil
	}

	parsedIDs := make([]uint64, len(tagIDs))
	for i, id := range tagIDs {
		parsedIDs[i] = ports.TagParseID(id.String())
	}

	var models []models.Proxy
	// Find proxies that have ALL specified tags
	err := r.db.WithContext(ctx).
		Preload("Tags").
		Joins("JOIN proxy_tags ON proxy_tags.proxy_id = proxies.id").
		Where("proxy_tags.tag_id IN ?", parsedIDs).
		Group("proxies.id").
		Having("COUNT(DISTINCT proxy_tags.tag_id) = ?", len(parsedIDs)).
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	proxies := make([]*proxy.Proxy, len(models))
	for i, model := range models {
		p, err := ports.ProxyPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		proxies[i] = p
	}

	return proxies, nil
}

func (r *ProxyRepository) FindActive(ctx context.Context) ([]*proxy.Proxy, error) {
	return r.FindByStatus(ctx, proxy.ProxyStatusActive)
}

func (r *ProxyRepository) Delete(ctx context.Context, id proxy.ProxyID) error {
	return r.db.WithContext(ctx).Where("id = ?", ports.ProxyParseID(id.String())).Delete(&models.Proxy{}).Error
}

func (r *ProxyRepository) Exists(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Proxy{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}
