package scenario

import (
	"errors"
)

type Context struct {
	Blocks []Node
	Edges  []Edge
}

func NewContext(blocks []Node, edges []Edge) Context {
	return Context{
		Blocks: blocks,
		Edges:  edges,
	}
}

func (c Context) IsEmpty() bool {
	return len(c.Blocks) == 0 && len(c.Edges) == 0
}

type Node struct {
	Id       string
	NodeType string
	Position Point2D
}

func NewNode(id, nodeType string, position Point2D) (Node, error) {
	if id == "" {
		return Node{}, errors.New("node id cannot be empty")
	}
	if nodeType == "" {
		return Node{}, errors.New("node type cannot be empty")
	}
	return Node{
		Id:       id,
		NodeType: nodeType,
		Position: position,
	}, nil
}

type Edge struct {
	Id           string
	Source       string
	Target       string
	SourceHandle string
	TargetHandle string
	Condition    string
}

func NewEdge(id, source, target, sourceHandle, targetHandle string) (Edge, error) {
	if id == "" {
		return Edge{}, errors.New("edge id cannot be empty")
	}
	if source == "" {
		return Edge{}, errors.New("edge source cannot be empty")
	}
	if target == "" {
		return Edge{}, errors.New("edge target cannot be empty")
	}
	return Edge{
		Id:           id,
		Source:       source,
		Target:       target,
		SourceHandle: sourceHandle,
		TargetHandle: targetHandle,
	}, nil
}

type Point2D struct {
	X float32
	Y float32
}

func NewPoint2D(x, y float32) Point2D {
	return Point2D{X: x, Y: y}
}

type InputData struct {
	Parameters []NodeParameters
}

func NewInputData(parameters []NodeParameters) InputData {
	return InputData{Parameters: parameters}
}

func (id InputData) IsEmpty() bool {
	return len(id.Parameters) == 0
}

type NodeParameters struct {
	BlockID string
	Input   []Parameter
	Output  []Parameter
}

func NewNodeParameters(blockID string, input, output []Parameter) (NodeParameters, error) {
	if blockID == "" {
		return NodeParameters{}, errors.New("block id cannot be empty")
	}
	return NodeParameters{
		BlockID: blockID,
		Input:   input,
		Output:  output,
	}, nil
}

type Parameter struct {
	Name  string
	Value interface{}
}

func NewParameter(name string, value interface{}) (Parameter, error) {
	if name == "" {
		return Parameter{}, errors.New("parameter name cannot be empty")
	}
	return Parameter{
		Name:  name,
		Value: value,
	}, nil
}

// Parameters represents scenario parameters
type Parameters struct {
	Input  []ParameterItem
	Output []ParameterItem
}

func NewParameters(input, output []ParameterItem) Parameters {
	return Parameters{
		Input:  input,
		Output: output,
	}
}

func (p Parameters) IsEmpty() bool {
	return len(p.Input) == 0 && len(p.Output) == 0
}

type ParameterItem struct {
	Parameter Parameter
	ParamType string
	Values    []string
}

func NewParameterItem(parameter Parameter, paramType string, values []string) ParameterItem {
	return ParameterItem{
		Parameter: parameter,
		ParamType: paramType,
		Values:    values,
	}
}
