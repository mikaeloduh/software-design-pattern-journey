package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSingleComparator(t *testing.T) {
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
			"Diamonds-Jack is greater than Spades-2",
			fields{nil},
			args{
				topCards:  NewSinglePattern([]BigTwoCard{{Suit: Diamonds, Rank: Jack}}),
				playCards: NewSinglePattern([]BigTwoCard{{Suit: Spades, Rank: Two}}),
			},
			true,
		},
		{
			"Diamonds-8 is not greater then Hearts-7",
			fields{nil},
			args{
				topCards:  NewSinglePattern([]BigTwoCard{{Suit: Diamonds, Rank: Eight}}),
				playCards: NewSinglePattern([]BigTwoCard{{Suit: Hearts, Rank: Seven}}),
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := SinglePatternComparator{Next: tt.fields.Next}
			got := h.Do(tt.args.topCards, tt.args.playCards)

			assert.Equal(t, tt.want, got)
		})
	}
}
