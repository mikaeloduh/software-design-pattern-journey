package service

import "github.com/dominikbraun/graph"

type RelationshipGraph graph.Graph[string, string]

type IRelationshipAnalyzer interface {
	Parse(script string)
	GetMutualFriends(name1, name2 string) []string
}
