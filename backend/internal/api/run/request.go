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

type GetRunByIDRequest struct {
	ID uint `path:"id" required:"true"`
}

type RunCommandRequest struct {
	GetRunByIDRequest
	Command string `json:"command"` // could be "start" | "stop" | "pause" | "resume" | "restart"
}
