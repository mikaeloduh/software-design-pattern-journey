package service

import "github.com/dominikbraun/graph"

type RelationshipGraph struct {
	graph.Graph[string, string]
}

func (g *RelationshipGraph) HasConnection(name1, name2 string) bool {
	return false
}
