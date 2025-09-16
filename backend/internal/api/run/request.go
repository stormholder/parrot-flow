package run

import "parrotflow/internal/api"

type RunQuery struct {
	Name    string   `json:"name,omitempty"`
	State   []string `json:"state,omitempty"`
	OS      string   `json:"os,omitempty"`
	Browser string   `json:"browser,omitempty"`
	api.PageQuery
	api.OrderByQuery
}
