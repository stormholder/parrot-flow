package scenario

type Point2D struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Node struct {
	ID       string  `json:"id"`
	Type     string  `json:"type"`
	Position Point2D `json:"position"`
}

type Edge struct {
	ID           string `json:"id"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	SourceHandle string `json:"sourceHandle"`
	TargetHandle string `json:"targetHandle"`
	Condition    string `json:"condition,omitempty"`
}

type ScenarioContext struct {
	Blocks []Node `json:"blocks"`
	Edges  []Edge `json:"edges"`
}

type ScenarioBlockParameter struct {
	Name  string `json:"name"`
	Value any    `json:"value"`
}

type ScenarioParameterItem struct {
	ScenarioBlockParameter
	Type   string    `json:"type"`
	Values *[]string `json:"values"`
}

type ScenarioParameters struct {
	Input  []ScenarioParameterItem `json:"input"`
	Output []ScenarioParameterItem `json:"output"`
}

type ScenarioNodeParameters struct {
	BlockID string                   `json:"blockId"`
	Input   []ScenarioBlockParameter `json:"input"`
	Output  []ScenarioBlockParameter `json:"output"`
}

type ScenarioInputData []ScenarioNodeParameters

type ScenarioPatch struct {
	Name        *string                 `json:"name,omitempty"`
	Description *string                 `json:"description,omitempty"`
	Tag         *string                 `json:"tag,omitempty"`
	Icon        *string                 `json:"icon,omitempty"`
	Context     *ScenarioContext        `json:"context,omitempty"`
	InputData   *interface{}            `json:"input_data,omitempty"`
	Parameters  *ScenarioNodeParameters `json:"parameters,omitempty"`
}
