package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFullHouseComparator(t *testing.T) {
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
			"Happy test",
			fields{nil},
			args{
				topCards: NewFullHousePattern([]BigTwoCard{
					{Suit: Spades, Rank: Three},
					{Suit: Diamonds, Rank: Three},
					{Suit: Hearts, Rank: Three},
					{Suit: Spades, Rank: Seven},
					{Suit: Hearts, Rank: Seven},
				}),
				playCards: NewFullHousePattern([]BigTwoCard{
					{Suit: Diamonds, Rank: Four},
					{Suit: Hearts, Rank: Four},
					{Suit: Spades, Rank: Four},
					{Suit: Diamonds, Rank: Seven},
					{Suit: Clubs, Rank: Seven},
				}),
			},
			true,
		},
		{
			"Happy test",
			fields{nil},
			args{
				topCards: NewFullHousePattern([]BigTwoCard{
					{Suit: Diamonds, Rank: Five},
					{Suit: Hearts, Rank: Five},
					{Suit: Spades, Rank: Five},
					{Suit: Diamonds, Rank: Two},
					{Suit: Clubs, Rank: Two},
				}),
				playCards: NewFullHousePattern([]BigTwoCard{
					{Suit: Spades, Rank: Ace},
					{Suit: Diamonds, Rank: Ace},
					{Suit: Hearts, Rank: Ace},
					{Suit: Spades, Rank: Three},
					{Suit: Hearts, Rank: Three},
				}),
			},
			true,
		},
		{
			"Unhappy test",
			fields{nil},
			args{
				topCards: NewFullHousePattern([]BigTwoCard{
					{Suit: Spades, Rank: Two},
					{Suit: Diamonds, Rank: Two},
					{Suit: Hearts, Rank: Two},
					{Suit: Spades, Rank: Three},
					{Suit: Hearts, Rank: Three},
				}),
				playCards: NewFullHousePattern([]BigTwoCard{
					{Suit: Diamonds, Rank: Three},
					{Suit: Hearts, Rank: Three},
					{Suit: Spades, Rank: Three},
					{Suit: Diamonds, Rank: Four},
					{Suit: Clubs, Rank: Four},
				}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := FullHousePatternComparator{Next: tt.fields.Next}
			got := p.Do(tt.args.topCards, tt.args.playCards)

			assert.Equalf(t, tt.want, got, "Do(%v, %v)", tt.args.topCards, tt.args.playCards)
		})
	}
}
