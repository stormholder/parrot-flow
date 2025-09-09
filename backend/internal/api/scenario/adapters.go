package scenario

import (
	"net/url"
	"parrotflow/internal/api"
	"parrotflow/internal/models"
	"strings"
)

func GetScenarioQuery(params url.Values) ScenarioQuery {
	generic := api.ParseGenericQuery(params)
	var query ScenarioQuery = ScenarioQuery{
		PageQuery:    generic.PageQuery,
		OrderByQuery: generic.OrderByQuery,
	}

	name := params.Get("name")
	rawTags := params.Get("tag")
	if name != "" {
		query.Name = name
	}
	if rawTags != "" {
		query.Tags = strings.Split(rawTags, ",")
	}
	return query
}

func GetScenarioListResponse(pages api.Pages) api.Pages {
	var list []models.ScenarioBase = make([]models.ScenarioBase, len(pages.Data.([]models.Scenario)))
	for i, s := range pages.Data.([]models.Scenario) {
		list[i] = models.ScenarioBase{
			Model:       s.Model,
			Name:        s.Name,
			Description: s.Description,
			Tag:         s.Tag,
			Icon:        s.Icon,
		}
	}
	return api.Pages{
		CurrentPage: pages.CurrentPage,
		Total:       pages.Total,
		PerPage:     pages.PerPage,
		Pages:       pages.Pages,
		Data:        list,
	}
}
