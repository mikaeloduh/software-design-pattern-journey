package service

import "github.com/dominikbraun/graph"

type RelationshipGraph graph.Graph[string, string]

type IRelationshipAnalyzer interface {
	Parse(script string)
	GetMutualFriends(name1, name2 string) []string
}

type RelationshipAnalyzerAdaptor struct {
	superRelationshipAnalyzer SuperRelationshipAnalyzer
	RelationshipGraph         graph.Graph[string, string]
}

func (a *RelationshipAnalyzerAdaptor) Parse(script string) {
}

func (a *RelationshipAnalyzerAdaptor) GetMutualFriends(name1, name2 string) []string {
	mutualFriends := make([]string, 0)

	_ = graph.DFS[string, string](a.superRelationshipAnalyzer.RelationshipGraph, name1, func(value string) bool {
		if a.superRelationshipAnalyzer.IsMutualFriend(value, name1, name2) && value != name1 && value != name2 {
			mutualFriends = append(mutualFriends, value)
		}
		return false
	})

	return mutualFriends
}
