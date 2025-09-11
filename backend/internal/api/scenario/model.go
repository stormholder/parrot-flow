package scenario

type Point2D struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

type Node struct {
	ID       string  `json:"id"`
	Position Point2D `json:"position"`
}

type Edge struct {
	ID           string `json:"id"`
	Source       string `json:"source"`
	Target       string `json:"target"`
	SourceHandle string `json:"sourceHandle"`
	TargetHandle string `json:"targetHandle"`
}

type ScenarioPayload struct {
	Blocks []Node `json:"blocks"`
	Edges  []Edge `json:"edges"`
}

type ScenarioPatch struct {
	Name        *string          `json:"name,omitempty"`
	Description *string          `json:"description,omitempty"`
	Tag         *string          `json:"tag,omitempty"`
	Icon        *string          `json:"icon,omitempty"`
	Payload     *ScenarioPayload `json:"payload,omitempty"`
	InputData   *interface{}     `json:"input_data,omitempty"`
	Parameters  *interface{}     `json:"parameters,omitempty"`
}
