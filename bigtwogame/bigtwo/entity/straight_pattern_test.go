package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStraightComparator(t *testing.T) {
	type fields struct {
		Next IPatternComparator
	}
	type args struct {
		topCards  ICardPattern
		playCards ICardPattern
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"happy test",
			fields{nil},
			args{
				topCards: NewStraightPattern([]BigTwoCard{
					{Suit: Hearts, Rank: Three},
					{Suit: Spades, Rank: Four},
					{Suit: Diamonds, Rank: Five},
					{Suit: Clubs, Rank: Six},
					{Suit: Diamonds, Rank: Seven},
				}),
				playCards: NewStraightPattern([]BigTwoCard{
					{Suit: Spades, Rank: Six},
					{Suit: Diamonds, Rank: Seven},
					{Suit: Hearts, Rank: Eight},
					{Suit: Spades, Rank: Nine},
					{Suit: Hearts, Rank: Ten},
				}),
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := StraightComparator{Next: tt.fields.Next}
			got := p.Do(tt.args.topCards, tt.args.playCards)

			assert.Equalf(t, tt.want, got, "Do(%v, %v)", tt.args.topCards, tt.args.playCards)
		})
	}
}
