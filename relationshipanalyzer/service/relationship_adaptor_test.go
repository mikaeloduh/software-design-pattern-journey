package service

import (
	"fmt"
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
			want: make([]string, 0),
		},
		{
			name: "test get all of B and C's mutual friends",
			args: args{
				name1: "C",
				name2: "B",
			},
			want: make([]string, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &RelationshipAnalyzerAdaptor{
				superRelationshipAnalyzer: FakeSuperRelationshipAnalyzer(),
			}

			a.Parse("")
			got := a.GetMutualFriends(tt.args.name1, tt.args.name2)

			fmt.Printf("%s and %s's mutual friends: %v\n", tt.args.name1, tt.args.name2, got)
		})
	}
}

func FakeSuperRelationshipAnalyzer() SuperRelationshipAnalyzer {
	return SuperRelationshipAnalyzer{SuperRelationship: testGraph()}
}
