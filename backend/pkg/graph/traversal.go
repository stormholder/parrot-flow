package graph

import "errors"

func NewGraph(vertices int) *Graph {
	return &Graph{
		Vertices:    vertices,
		Inputs:      make(map[string][]string),
		Outputs:     make(map[string][]string),
		BranchMarks: make(map[string][]string),
	}
}

func (g *Graph) AddEdge(source, dest string, handle string) error {
	g.Outputs[source] = append(g.Outputs[source], dest)
	g.Inputs[dest] = append(g.Inputs[dest], source)
	if handle != "" {
		g.BranchMarks[dest] = append(g.BranchMarks[dest], handle)
	}
	return nil
}

func KeyExists(m map[string][]string, key string) bool {
	_, exists := m[key]
	return exists
}

func invertSlice(slice []string) {
	n := len(slice)

	// Check if the slice has less than 2 elements and return early
	if n < 2 {
		return
	}

	for i := 0; i < (n / 2); i++ {
		// Swap the elements at index `i` and `n-1-i`
		slice[i], slice[n-1-i] = slice[n-1-i], slice[i]
	}
}

func (g *Graph) Dfs() ([]string, error) {
	var (
		acyclic       = true
		order         []string
		permanentMark = make(map[string]bool)
		temporaryMark = make(map[string]bool)
		visit         func(string)
	)

	visit = func(u string) {
		if temporaryMark[u] {
			acyclic = false
		} else if !(temporaryMark[u] || permanentMark[u]) {
			temporaryMark[u] = true
			for _, v := range g.Inputs[u] {
				visit(v)
				marks, ok := g.BranchMarks[v]
				if ok {
					g.BranchMarks[u] = append(g.BranchMarks[u], marks...)
				}
				if !acyclic {
					return
				}
			}
			delete(temporaryMark, u)
			permanentMark[u] = true
			order = append([]string{u}, order...)
		}
	}

	for u := range g.Inputs {
		if !permanentMark[u] {
			visit(u)
			if !acyclic {
				return order, errors.New("not a DAG")
			}
		}
	}
	invertSlice(order)
	return order, nil
}

func (g *Graph) Bfs(node string, visited map[string]bool) []string {
	next := []string{}

	for _, v := range g.Outputs[node] {
		if !visited[v] {
			next = append(next, v)
		}
	}

	return next
}
