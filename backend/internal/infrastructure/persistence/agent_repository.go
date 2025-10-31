package persistence

import (
	"context"
	"errors"
	"parrotflow/internal/domain/agent"
	"parrotflow/internal/domain/tag"
	"parrotflow/internal/models"
	"parrotflow/internal/ports"
	"time"

	"gorm.io/gorm"
)

type AgentRepository struct {
	db *gorm.DB
}

func NewAgentRepository(db *gorm.DB) *AgentRepository {
	return &AgentRepository{db: db}
}

func (r *AgentRepository) Save(ctx context.Context, a *agent.Agent) error {
	model, err := ports.AgentDomainEntityToPersistence(a)
	if err != nil {
		return err
	}

	// Load tags from IDs
	if len(a.Tags) > 0 {
		tagIDs := make([]uint64, len(a.Tags))
		for i, tagID := range a.Tags {
			tagIDs[i] = ports.TagParseID(tagID.String())
		}
		var tags []models.Tag
		if err := r.db.WithContext(ctx).Where("id IN ?", tagIDs).Find(&tags).Error; err != nil {
			return err
		}
		model.Tags = tags
	}

	// Save agent
	if err = r.db.WithContext(ctx).Save(model).Error; err != nil {
		return err
	}

	// Update tag associations
	if err = r.db.WithContext(ctx).Model(model).Association("Tags").Replace(model.Tags); err != nil {
		return err
	}

	return nil
}

func (r *AgentRepository) FindByID(ctx context.Context, id agent.AgentID) (*agent.Agent, error) {
	var model models.Agent
	if err := r.db.WithContext(ctx).Preload("Tags").Where("id = ?", ports.AgentParseID(id.String())).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("agent not found")
		}
		return nil, err
	}

	return ports.AgentPersistenceToDomainEntity(&model)
}

func (r *AgentRepository) FindByName(ctx context.Context, name string) (*agent.Agent, error) {
	var model models.Agent
	if err := r.db.WithContext(ctx).Preload("Tags").Where("name = ?", name).First(&model).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, errors.New("agent not found")
		}
		return nil, err
	}

	return ports.AgentPersistenceToDomainEntity(&model)
}

func (r *AgentRepository) FindAll(ctx context.Context) ([]*agent.Agent, error) {
	var models []models.Agent
	if err := r.db.WithContext(ctx).Preload("Tags").Order("name ASC").Find(&models).Error; err != nil {
		return nil, err
	}

	agents := make([]*agent.Agent, len(models))
	for i, model := range models {
		a, err := ports.AgentPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		agents[i] = a
	}

	return agents, nil
}

// FindByCriteria retrieves agents matching the specified criteria
// Supports combining multiple filters (status AND tags AND browser, etc.)
func (r *AgentRepository) FindByCriteria(ctx context.Context, criteria agent.SearchCriteria) ([]*agent.Agent, error) {
	query := r.db.WithContext(ctx).Preload("Tags")

	// Apply status filter if specified
	if criteria.Status != nil {
		query = query.Where("status = ?", criteria.Status.String())
	}

	// Apply tag filter if specified
	if len(criteria.TagIDs) > 0 {
		tagIDs := make([]uint64, len(criteria.TagIDs))
		for i, tagID := range criteria.TagIDs {
			tagIDs[i] = ports.TagParseID(tagID.String())
		}
		// Find agents that have ALL specified tags
		query = query.Joins("JOIN agent_tags ON agent_tags.agent_id = agents.id").
			Where("agent_tags.tag_id IN ?", tagIDs).
			Group("agents.id").
			Having("COUNT(DISTINCT agent_tags.tag_id) = ?", len(tagIDs))
	}

	// Apply browser type filter if specified
	if criteria.BrowserType != nil {
		query = query.Where("capabilities::jsonb @> ?", `{"browsers":[{"type":"`+criteria.BrowserType.String()+`"}]}`)
	}

	// Apply platform filter if specified
	if criteria.Platform != nil {
		query = query.Where("capabilities::jsonb @> ?", `{"os":{"platform":"`+criteria.Platform.String()+`"}}`)
	}

	var models []models.Agent
	if err := query.Find(&models).Error; err != nil {
		return nil, err
	}

	// Convert to domain entities
	agents, err := ConvertSliceToDomainPtr(models, ports.AgentPersistenceToDomainEntity)
	if err != nil {
		return nil, err
	}

	// Apply health filter if requested (in-memory filtering)
	if criteria.OnlyHealthy && criteria.HeartbeatTimeout > 0 {
		healthyAgents := make([]*agent.Agent, 0, len(agents))
		for _, a := range agents {
			if a.IsHealthy(criteria.HeartbeatTimeout) {
				healthyAgents = append(healthyAgents, a)
			}
		}
		return healthyAgents, nil
	}

	return agents, nil
}

func (r *AgentRepository) FindByStatus(ctx context.Context, status agent.AgentStatus) ([]*agent.Agent, error) {
	var models []models.Agent
	if err := r.db.WithContext(ctx).Preload("Tags").Where("status = ?", status.String()).Find(&models).Error; err != nil {
		return nil, err
	}

	agents := make([]*agent.Agent, len(models))
	for i, model := range models {
		a, err := ports.AgentPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		agents[i] = a
	}

	return agents, nil
}

func (r *AgentRepository) FindByTags(ctx context.Context, tagIDs []tag.TagID) ([]*agent.Agent, error) {
	if len(tagIDs) == 0 {
		return []*agent.Agent{}, nil
	}

	parsedIDs := make([]uint64, len(tagIDs))
	for i, id := range tagIDs {
		parsedIDs[i] = ports.TagParseID(id.String())
	}

	var models []models.Agent
	// Find agents that have ALL specified tags
	err := r.db.WithContext(ctx).
		Preload("Tags").
		Joins("JOIN agent_tags ON agent_tags.agent_id = agents.id").
		Where("agent_tags.tag_id IN ?", parsedIDs).
		Group("agents.id").
		Having("COUNT(DISTINCT agent_tags.tag_id) = ?", len(parsedIDs)).
		Find(&models).Error

	if err != nil {
		return nil, err
	}

	agents := make([]*agent.Agent, len(models))
	for i, model := range models {
		a, err := ports.AgentPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		agents[i] = a
	}

	return agents, nil
}

func (r *AgentRepository) FindAvailable(ctx context.Context) ([]*agent.Agent, error) {
	var models []models.Agent
	// Available agents: status is online/idle/busy (not offline/disconnected)
	// and they have capacity (current_run_count < max from capabilities)
	// Note: We can't easily query JSON in GORM, so we'll filter in application code
	if err := r.db.WithContext(ctx).
		Preload("Tags").
		Where("status IN ?", []string{"online", "idle", "busy"}).
		Find(&models).Error; err != nil {
		return nil, err
	}

	agents := make([]*agent.Agent, 0, len(models))
	for _, model := range models {
		a, err := ports.AgentPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		// Filter by capacity
		if a.CanAcceptRun() {
			agents = append(agents, a)
		}
	}

	return agents, nil
}

func (r *AgentRepository) FindByBrowserType(ctx context.Context, browserType agent.BrowserType) ([]*agent.Agent, error) {
	// This requires querying JSON, which is database-specific
	// For PostgreSQL, we can use jsonb operators
	var models []models.Agent
	if err := r.db.WithContext(ctx).
		Preload("Tags").
		Where("capabilities::jsonb @> ?", `{"browsers":[{"type":"`+browserType.String()+`"}]}`).
		Find(&models).Error; err != nil {
		return nil, err
	}

	agents := make([]*agent.Agent, len(models))
	for i, model := range models {
		a, err := ports.AgentPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		agents[i] = a
	}

	return agents, nil
}

func (r *AgentRepository) FindByPlatform(ctx context.Context, platform agent.Platform) ([]*agent.Agent, error) {
	// Query JSON for platform
	var models []models.Agent
	if err := r.db.WithContext(ctx).
		Preload("Tags").
		Where("capabilities::jsonb @> ?", `{"os":{"platform":"`+platform.String()+`"}}`).
		Find(&models).Error; err != nil {
		return nil, err
	}

	agents := make([]*agent.Agent, len(models))
	for i, model := range models {
		a, err := ports.AgentPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		agents[i] = a
	}

	return agents, nil
}

func (r *AgentRepository) FindStaleAgents(ctx context.Context, heartbeatTimeout time.Duration) ([]*agent.Agent, error) {
	cutoffTime := time.Now().Add(-heartbeatTimeout)

	var models []models.Agent
	if err := r.db.WithContext(ctx).
		Preload("Tags").
		Where("last_heartbeat_at < ?", cutoffTime).
		Where("status NOT IN ?", []string{"offline", "disconnected"}).
		Find(&models).Error; err != nil {
		return nil, err
	}

	agents := make([]*agent.Agent, len(models))
	for i, model := range models {
		a, err := ports.AgentPersistenceToDomainEntity(&model)
		if err != nil {
			return nil, err
		}
		agents[i] = a
	}

	return agents, nil
}

func (r *AgentRepository) Delete(ctx context.Context, id agent.AgentID) error {
	return r.db.WithContext(ctx).Where("id = ?", ports.AgentParseID(id.String())).Delete(&models.Agent{}).Error
}

func (r *AgentRepository) Exists(ctx context.Context, id agent.AgentID) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Agent{}).Where("id = ?", ports.AgentParseID(id.String())).Count(&count).Error
	return count > 0, err
}

func (r *AgentRepository) ExistsByName(ctx context.Context, name string) (bool, error) {
	var count int64
	err := r.db.WithContext(ctx).Model(&models.Agent{}).Where("name = ?", name).Count(&count).Error
	return count > 0, err
}
