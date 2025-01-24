package service

import (
	"testing"

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
		t.Run("unit "+tt.name, func(t *testing.T) {
			a := FakeNewSuperRelationshipAnalyzer()

			got := a.IsMutualFriend(tt.args.target, tt.args.name2, tt.args.name3)

			assert.Equal(t, tt.want, got)
		})
	}

	for _, tt := range tests {
		t.Run("integration "+tt.name, func(t *testing.T) {
			a := NewSuperRelationshipAnalyzer()
			a.Init(testSuperScript)

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
		{
			name: "test A and E is not friend",
			args: args{
				name1: "A",
				name2: "E",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run("unit "+tt.name, func(t *testing.T) {
			a := FakeNewSuperRelationshipAnalyzer()

			got := a.isFriend(tt.args.name1, tt.args.name2)

			assert.Equal(t, tt.want, got)
		})
	}

	for _, tt := range tests {
		t.Run("integration "+tt.name, func(t *testing.T) {
			a := NewSuperRelationshipAnalyzer()
			a.Init(testSuperScript)

			got := a.isFriend(tt.args.name1, tt.args.name2)

			assert.Equal(t, tt.want, got)
		})
	}
}

func FakeNewSuperRelationshipAnalyzer() *SuperRelationshipAnalyzer {
	return &SuperRelationshipAnalyzer{superRelationshipGraph: testGraph()}
}
