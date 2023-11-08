package service

import (
	"github.com/dominikbraun/graph"
)

type SuperRelationshipAnalyzer struct {
	NameGraph graph.Graph[string, string]
}

func (a *SuperRelationshipAnalyzer) Init(script string) {

}

func (a *SuperRelationshipAnalyzer) IsMutualFriend(target, name2, name3 string) bool {

	return a.isFriend(target, name2) && a.isFriend(target, name3)
}

func (a *SuperRelationshipAnalyzer) isFriend(name1, name2 string) bool {
	isFound := false

	_ = graph.DFS[string, string](a.NameGraph, name1, func(value string) bool {
		if value == name2 {
			isFound = true
			return true
		}
		return false
	})

	return isFound
}
