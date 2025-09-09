package graph

type Graph struct {
	Vertices    int
	Inputs      map[string][]string
	Outputs     map[string][]string
	BranchMarks map[string][]string
}
