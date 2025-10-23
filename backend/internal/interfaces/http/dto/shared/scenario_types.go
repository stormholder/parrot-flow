package shared

// Scenario value object DTOs - shared across queries and commands

type ContextDTO struct {
	Blocks []NodeDTO `json:"blocks"`
	Edges  []EdgeDTO `json:"edges"`
}

type NodeDTO struct {
	ID       string    `json:"id"`
	NodeType string    `json:"node_type"`
	Position Point2D   `json:"position"`
}

type EdgeDTO struct {
	ID           string `json:"id"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	SourceHandle string `json:"source_handle"`
	TargetHandle string `json:"target_handle"`
	Condition    string `json:"condition,omitempty"`
}

type Point2D struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type InputDataDTO struct {
	Parameters []NodeParametersDTO `json:"parameters"`
}

type NodeParametersDTO struct {
	BlockID string         `json:"block_id"`
	Input   []ParameterDTO `json:"input"`
	Output  []ParameterDTO `json:"output"`
}

type ParameterDTO struct {
	Name  string      `json:"name"`
	Value interface{} `json:"value"`
}

type ParametersDTO struct {
	Input  []ParameterItemDTO `json:"input"`
	Output []ParameterItemDTO `json:"output"`
}

type ParameterItemDTO struct {
	Parameter ParameterDTO `json:"parameter"`
	ParamType string       `json:"param_type"`
	Values    []string     `json:"values"`
}
