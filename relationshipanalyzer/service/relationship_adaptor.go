package service

import (
	"fmt"
	"strings"
)

type RelationshipAnalyzerAdaptor struct {
	superRelationshipAnalyzer *SuperRelationshipAnalyzer
	names                     []string
	relationshipGraph         RelationshipGraph
}

func NewRelationshipAnalyzerAdaptor() *RelationshipAnalyzerAdaptor {
	return &RelationshipAnalyzerAdaptor{
		superRelationshipAnalyzer: NewSuperRelationshipAnalyzer(),
	}
}

func (a *RelationshipAnalyzerAdaptor) Parse(script string) {
	// Create a set to store unique elements
	elements := make(map[string]bool)

	// Create a map to store relationships
	relationships := make(map[string][]string)

	// Create a set to keep track of processed relationships
	processed := make(map[string]bool)

	// Split input into lines
	lines := strings.Split(script, "\n")

	// Process each line
	for _, line := range lines {
		// Split line into components
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			continue
		}

		// Trim spaces
		node := strings.TrimSpace(parts[0])
		neighbors := strings.Fields(parts[1])

		// Add relationships to the map and check for duplicates
		for _, neighbor := range neighbors {
			relationship := fmt.Sprintf("%s -- %s", node, neighbor)
			reverseRelationship := fmt.Sprintf("%s -- %s", neighbor, node)

			if !processed[relationship] && !processed[reverseRelationship] {
				relationships[node] = append(relationships[node], neighbor)
				processed[relationship] = true
				elements[node] = true
				elements[neighbor] = true
			}
		}

	}
	// Convert the set of elements to a list
	for element := range elements {
		a.names = append(a.names, element)
	}
	fmt.Println(a.names)

	superInput := ""
	// Print relationships in the desired format
	for node, neighbors := range relationships {
		for _, neighbor := range neighbors {
			superInput += fmt.Sprintf("%s -- %s\n", node, neighbor)
		}
	}

	a.superRelationshipAnalyzer.Init(superInput)
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

func (a *RelationshipAnalyzerAdaptor) HasConnection(name1, name2 string) bool {
	return a.relationshipGraph.HasConnection(name1, name2)
}

func filter(ss []string, test func(string) bool) (ret []string) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
