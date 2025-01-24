package service

import (
	"github.com/dominikbraun/graph"
	"strings"
)

type SuperRelationshipAnalyzer struct {
	superRelationshipGraph graph.Graph[string, string]
}

func NewSuperRelationshipAnalyzer() *SuperRelationshipAnalyzer {
	return &SuperRelationshipAnalyzer{}
}

func (a *SuperRelationshipAnalyzer) Init(script string) {
	// Create a new graph
	a.superRelationshipGraph = graph.New(graph.StringHash)

	// Split input into lines
	lines := strings.Split(script, "\n")

	// Process each line
	for _, line := range lines {
		// Split line into components
		parts := strings.Split(line, "--")
		if len(parts) != 2 {
			continue
		}

		// Trim spaces
		source := strings.TrimSpace(parts[0])
		destination := strings.TrimSpace(parts[1])

		// Add vertices and edges to the graph
		_ = a.superRelationshipGraph.AddVertex(source)
		_ = a.superRelationshipGraph.AddVertex(destination)
		_ = a.superRelationshipGraph.AddEdge(source, destination)
	}
}

func (a *SuperRelationshipAnalyzer) IsMutualFriend(target, name2, name3 string) bool {

	return a.isFriend(target, name2) && a.isFriend(target, name3)
}

func (a *SuperRelationshipAnalyzer) isFriend(name1, name2 string) bool {
	isFound := false

	_ = graph.BFSWithDepth[string, string](a.superRelationshipGraph, name1, func(value string, depth int) bool {
		if value == name2 {
			isFound = true
			return true
		}

		return depth > 1
	})

	return isFound
}
