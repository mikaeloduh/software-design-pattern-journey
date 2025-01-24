package service

import "github.com/dominikbraun/graph"

// IRelationshipGraph
type IRelationshipGraph interface {
	HasConnection(name1, name2 string) bool
}

// RelationshipGraph
type RelationshipGraph struct {
	graph.Graph[string, string]
}

func (g *RelationshipGraph) HasConnection(name1, name2 string) bool {
	panic("to be implement")
}

func (g *RelationshipGraph) GetElements() []string {
	panic("to be implement")
}
