package service

import (
	"testing"

	"github.com/dominikbraun/graph"
	"github.com/stretchr/testify/assert"
)

func TestSuperRelationshipAnalyzer_IsMutualFriend(t *testing.T) {

	type args struct {
		target string
		name2  string
		name3  string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test A is B, C's mutual fired",
			args: args{
				target: "A",
				name2:  "B",
				name3:  "C",
			},
			want: true,
		},
		{
			name: "test A is not B, F's mutual fired",
			args: args{
				target: "A",
				name2:  "B",
				name3:  "F",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SuperRelationshipAnalyzer{RelationshipGraph: testGraph()}

			got := a.IsMutualFriend(tt.args.target, tt.args.name2, tt.args.name3)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSuperRelationshipAnalyzer_isFriend(t *testing.T) {

	type args struct {
		name1 string
		name2 string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test A and B is friend",
			args: args{
				name1: "A",
				name2: "B",
			},
			want: true,
		},
		{
			name: "test A and C is friend",
			args: args{
				name1: "A",
				name2: "C",
			},
			want: true,
		},
		{
			name: "test A and F is not friend",
			args: args{
				name1: "A",
				name2: "F",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &SuperRelationshipAnalyzer{RelationshipGraph: testGraph()}

			got := a.isFriend(tt.args.name1, tt.args.name2)

			assert.Equal(t, tt.want, got)
		})
	}
}

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
	_ = g.AddEdge("D", "P")
	_ = g.AddEdge("E", "J")
	_ = g.AddEdge("E", "K")
	_ = g.AddEdge("E", "L")
	_ = g.AddEdge("F", "Z")

	return g
}
