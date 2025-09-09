package api

import (
	"net/url"

	"parrotflow/pkg/shared"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Pages struct {
	TotalItems  int         `json:"total"`
	TotalPages  int         `json:"pages"`
	RowsPerPage int         `json:"rpp"`
	CurrentPage int         `json:"page"`
	Data        interface{} `json:"data"`
}

type PageQuery struct {
	Page    int `json:"page,omitempty" query:"page"`
	PerPage int `json:"per_page,omitempty" query:"rpp"`
}

type OrderByQuery struct {
	Field     string `json:"field,omitempty" query:"order"`
	Direction string `json:"direction,omitempty" query:"dir"`
}

type GenericQuery struct {
	PageQuery
	OrderByQuery
}

func ParseGenericQuery(params url.Values) GenericQuery {
	var query GenericQuery = GenericQuery{}
	page := shared.ParseInt(params.Get("page"), 1)
	perPage := shared.ParseInt(params.Get("per_page"), 25)
	orderByField := shared.ParseString(params.Get("order_by"), "created_at")
	orderDirection := shared.ParseString(params.Get("order_by"), "desc")
	query.Page = page
	query.PerPage = perPage
	query.Field = orderByField
	query.Direction = orderDirection
	return query
}

func BuildQuery(ctx *gorm.DB, query GenericQuery, fields []string) *gorm.DB {
	if len(fields) > 0 {
		ctx = ctx.Select(fields)
	}

	if query.Field != "" {
		isDesc := false
		if query.Direction == "desc" {
			isDesc = true
		}
		ctx = ctx.Order(clause.OrderByColumn{Column: clause.Column{Name: query.Field}, Desc: isDesc})
	}

	ctx = ctx.Limit(query.PerPage)

	if query.Page > 1 {
		ctx = ctx.Offset(query.Page * query.PerPage)
	}

	return ctx
}
