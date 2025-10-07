package run

import (
	"context"
	"parrotflow/internal/domain/scenario"
)

type Repository interface {
	Save(ctx context.Context, run *Run) error
	FindByID(ctx context.Context, id RunID) (*Run, error)
	FindByScenarioID(ctx context.Context, scenarioID scenario.ScenarioID) ([]*Run, error)
	FindAll(ctx context.Context, criteria SearchCriteria) ([]*Run, error)
	Delete(ctx context.Context, id RunID) error
	Exists(ctx context.Context, id RunID) (bool, error)
}

type SearchCriteria struct {
	ScenarioID scenario.ScenarioID
	Status     string
	Limit      int
	Offset     int
	OrderBy    string
	OrderDir   string
}

func NewSearchCriteria() SearchCriteria {
	return SearchCriteria{
		Limit:    10,
		Offset:   0,
		OrderBy:  "created_at",
		OrderDir: "desc",
	}
}

func (sc SearchCriteria) WithScenarioID(scenarioID scenario.ScenarioID) SearchCriteria {
	sc.ScenarioID = scenarioID
	return sc
}

func (sc SearchCriteria) WithStatus(status string) SearchCriteria {
	sc.Status = status
	return sc
}

func (sc SearchCriteria) WithPagination(limit, offset int) SearchCriteria {
	sc.Limit = limit
	sc.Offset = offset
	return sc
}

func (sc SearchCriteria) WithOrdering(orderBy, orderDir string) SearchCriteria {
	sc.OrderBy = orderBy
	sc.OrderDir = orderDir
	return sc
}
