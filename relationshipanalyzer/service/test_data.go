package service

import "github.com/dominikbraun/graph"

var testScript string = `
A: B C D
B: A D E
C: A E G K M
D: A B K P
E: B C J K L
F: Z
`

var testSuperScript string = `
A -- B
A -- C
A -- D
B -- D
B -- E
C -- E
C -- G
C -- K
C -- M
D -- K
D -- P
E -- J
E -- K
E -- L
F -- Z
`

func testGraph() graph.Graph[string, string] {
	g := graph.New(graph.StringHash)

	_ = g.AddVertex("A")
	_ = g.AddVertex("B")
	_ = g.AddVertex("C")
	_ = g.AddVertex("D")
	_ = g.AddVertex("E")
	_ = g.AddVertex("F")
	_ = g.AddVertex("G")
	_ = g.AddVertex("K")
	_ = g.AddVertex("L")
	_ = g.AddVertex("M")
	_ = g.AddVertex("P")
	_ = g.AddVertex("J")
	_ = g.AddVertex("Z")

	_ = g.AddEdge("A", "B")
	_ = g.AddEdge("A", "C")
	_ = g.AddEdge("A", "D")
	_ = g.AddEdge("B", "D")
	_ = g.AddEdge("B", "E")
	_ = g.AddEdge("C", "E")
	_ = g.AddEdge("C", "G")
	_ = g.AddEdge("C", "K")
	_ = g.AddEdge("C", "M")
	_ = g.AddEdge("D", "K")
	_ = g.AddEdge("D", "P")
	_ = g.AddEdge("E", "J")
	_ = g.AddEdge("E", "K")
	_ = g.AddEdge("E", "L")
	_ = g.AddEdge("F", "Z")

	return g
}
