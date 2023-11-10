package service

import (
	"github.com/dominikbraun/graph"
)

type SuperRelationshipAnalyzer struct {
	SuperRelationship graph.Graph[string, string]
}

func (a *SuperRelationshipAnalyzer) Init(script string) {
}

func (a *SuperRelationshipAnalyzer) IsMutualFriend(target, name2, name3 string) bool {

	return a.isFriend(target, name2) && a.isFriend(target, name3)
}

func (a *SuperRelationshipAnalyzer) isFriend(name1, name2 string) bool {
	isFound := false

	_ = graph.BFSWithDepth[string, string](a.SuperRelationship, name1, func(value string, depth int) bool {
		if value == name2 {
			isFound = true
			return true
		}

		return depth > 1
	})

	return isFound
}
