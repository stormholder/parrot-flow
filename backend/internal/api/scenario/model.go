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
