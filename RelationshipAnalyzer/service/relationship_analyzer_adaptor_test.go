package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRelationshipAnalyzerAdaptor_GetMutualFriends(t *testing.T) {
	type args struct {
		name1 string
		name2 string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "test get all of A and B's mutual friends",
			args: args{
				name1: "B",
				name2: "A",
			},
			want: []string{"D"},
		},
		{
			name: "test get all of B and C's mutual friends",
			args: args{
				name1: "C",
				name2: "B",
			},
			want: []string{"A", "E"},
		},
	}
	for _, tt := range tests {
		t.Run("unit "+tt.name, func(t *testing.T) {
			a := FakeNewRelationshipAnalyzerAdaptor()

			got := a.GetMutualFriends(tt.args.name1, tt.args.name2)

			assert.ElementsMatch(t, tt.want, got)
		})
	}

	for _, tt := range tests {
		t.Run("integration "+tt.name, func(t *testing.T) {
			a := NewRelationshipAnalyzerAdaptor()
			a.Parse(testScript)

			got := a.GetMutualFriends(tt.args.name1, tt.args.name2)

			assert.ElementsMatch(t, tt.want, got)
		})
	}
}

func FakeNewRelationshipAnalyzerAdaptor() *RelationshipAnalyzerAdaptor {
	return &RelationshipAnalyzerAdaptor{
		superRelationshipAnalyzer: FakeNewSuperRelationshipAnalyzer(),
		friends:                   []string{"A", "B", "C", "D", "E", "F", "G", "J", "K", "M", "P", "L", "Z"},
	}
}
