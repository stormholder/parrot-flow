package scenario

import (
	"context"
)

type Repository interface {
	Save(ctx context.Context, scenario *Scenario) error
	FindByID(ctx context.Context, id ScenarioID) (*Scenario, error)
	FindByName(ctx context.Context, name string) (*Scenario, error)
	FindAll(ctx context.Context, criteria SearchCriteria) ([]*Scenario, error)
	Delete(ctx context.Context, id ScenarioID) error
	Exists(ctx context.Context, id ScenarioID) (bool, error)
}

type SearchCriteria struct {
	Name     string
	Tag      string
	Limit    int
	Offset   int
	OrderBy  string
	OrderDir string
}

func NewSearchCriteria() SearchCriteria {
	return SearchCriteria{
		Limit:    10,
		Offset:   0,
		OrderBy:  "created_at",
		OrderDir: "desc",
	}
}

func (sc SearchCriteria) WithName(name string) SearchCriteria {
	sc.Name = name
	return sc
}

func (sc SearchCriteria) WithTag(tag string) SearchCriteria {
	sc.Tag = tag
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
