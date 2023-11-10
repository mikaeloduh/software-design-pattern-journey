package service

import "github.com/dominikbraun/graph"

type RelationshipAnalyzerAdaptor struct {
	superRelationshipAnalyzer SuperRelationshipAnalyzer
	names                     []string
}

func (a *RelationshipAnalyzerAdaptor) Parse(script string) {
	//a.superRelationshipAnalyzer.Init()
	a.names = []string{"A", "B", "C", "D", "E", "F", "G", "J", "K", "M", "P", "L", "Z"}
}

func (a *RelationshipAnalyzerAdaptor) GetMutualFriends(name1, name2 string) []string {
	mutualFriends := make([]string, 0)

	searchCandidate := filter(a.names, func(s string) bool { return s != name1 && s != name2 })

	for _, tar := range searchCandidate {
		if a.superRelationshipAnalyzer.IsMutualFriend(tar, name1, name2) {
			mutualFriends = append(mutualFriends, tar)
		}
	}

	return mutualFriends
}

func (a *RelationshipAnalyzerAdaptor) GetMutualFriendsV0(name1, name2 string) []string {
	mutualFriends := make([]string, 0)

	_ = graph.DFS[string, string](a.superRelationshipAnalyzer.SuperRelationship, name1, func(value string) bool {
		if a.superRelationshipAnalyzer.IsMutualFriend(value, name1, name2) && value != name1 && value != name2 {
			mutualFriends = append(mutualFriends, value)
		}
		return false
	})

	return mutualFriends
}

func (a *RelationshipAnalyzerAdaptor) HasConnection(name1, name2 string) bool {
	return false
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
