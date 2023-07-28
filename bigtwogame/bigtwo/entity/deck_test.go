package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSinglePatternComparator(t *testing.T) {
	type fields struct {
		Next IPatternHandler
	}
	type args struct {
		topCards  []BigTwoCard
		playCards []BigTwoCard
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
				topCards:  []BigTwoCard{{Suit: Diamonds, Rank: Jack}},
				playCards: []BigTwoCard{{Suit: Spades, Rank: Two}},
			},
			true,
		},
		{
			"Diamonds-8 is not greater then Hearts-7",
			fields{nil},
			args{
				topCards:  []BigTwoCard{{Suit: Diamonds, Rank: Eight}},
				playCards: []BigTwoCard{{Suit: Hearts, Rank: Seven}},
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

func TestSinglePatternValidator(t *testing.T) {
	type fields struct {
		Next IPatternValidator
	}
	type args struct {
		cards []BigTwoCard
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"test Happy",
			fields{nil},
			args{[]BigTwoCard{{Suit: Hearts, Rank: Two}}},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := SinglePatternValidator{Next: tt.fields.Next}
			got := v.Do(tt.args.cards)

			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStraightPatternValidator(t *testing.T) {
	type fields struct {
		Next IPatternValidator
	}
	type args struct {
		cards []BigTwoCard
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"should be a straight",
			fields{nil},
			args{cards: []BigTwoCard{
				{Suit: Hearts, Rank: Three},
				{Suit: Diamonds, Rank: Four},
				{Suit: Clubs, Rank: Five},
				{Suit: Clubs, Rank: Six},
				{Suit: Diamonds, Rank: Seven}}},
			true,
		},
		{
			"should be not a straight",
			fields{nil},
			args{cards: []BigTwoCard{
				{Suit: Hearts, Rank: Two},
				{Suit: Diamonds, Rank: Four},
				{Suit: Clubs, Rank: Five},
				{Suit: Clubs, Rank: Six},
				{Suit: Diamonds, Rank: Seven}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := StraightPatternValidator{Next: tt.fields.Next}
			got := v.Do(tt.args.cards)

			assert.Equalf(t, tt.want, got, "Do(%v)", tt.args.cards)
		})
	}
}

func TestStraightPatternComparator(t *testing.T) {
	type fields struct {
		Next IPatternHandler
	}
	type args struct {
		topCards  []BigTwoCard
		playCards []BigTwoCard
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
				topCards: []BigTwoCard{
					{Suit: Hearts, Rank: Three},
					{Suit: Spades, Rank: Four},
					{Suit: Diamonds, Rank: Five},
					{Suit: Clubs, Rank: Six},
					{Suit: Diamonds, Rank: Seven},
				},
				playCards: []BigTwoCard{
					{Suit: Spades, Rank: Six},
					{Suit: Diamonds, Rank: Seven},
					{Suit: Hearts, Rank: Eight},
					{Suit: Spades, Rank: Nine},
					{Suit: Hearts, Rank: Ten},
				},
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := StraightPatternComparator{Next: tt.fields.Next}
			got := p.Do(tt.args.topCards, tt.args.playCards)

			assert.Equalf(t, tt.want, got, "Do(%v, %v)", tt.args.topCards, tt.args.playCards)
		})
	}
}

func TestFullHousePatternValidator(t *testing.T) {
	type fields struct {
		Next IPatternValidator
	}
	type args struct {
		cards []BigTwoCard
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"should be a full-house",
			fields{nil},
			args{[]BigTwoCard{
				{Suit: Diamonds, Rank: Three},
				{Suit: Clubs, Rank: Three},
				{Suit: Spades, Rank: Three},
				{Suit: Clubs, Rank: Two},
				{Suit: Hearts, Rank: Two}}},
			true,
		},
		{
			"should not be a full-house",
			fields{nil},
			args{[]BigTwoCard{
				{Suit: Diamonds, Rank: Three},
				{Suit: Clubs, Rank: Three},
				{Suit: Spades, Rank: Three},
				{Suit: Clubs, Rank: Two},
				{Suit: Hearts, Rank: Ace}}},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := FullHousePatternValidator{Next: tt.fields.Next}
			got := v.Do(tt.args.cards)

			assert.Equalf(t, tt.want, got, "Do(%v)", tt.args.cards)
		})
	}
}

func TestFullHousePatternComparator(t *testing.T) {
	type fields struct {
		Next IPatternHandler
	}
	type args struct {
		topCards  []BigTwoCard
		playCards []BigTwoCard
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
				topCards: []BigTwoCard{
					{Suit: Spades, Rank: Three},
					{Suit: Diamonds, Rank: Three},
					{Suit: Hearts, Rank: Three},
					{Suit: Spades, Rank: Seven},
					{Suit: Hearts, Rank: Seven},
				},
				playCards: []BigTwoCard{
					{Suit: Diamonds, Rank: Four},
					{Suit: Hearts, Rank: Four},
					{Suit: Spades, Rank: Four},
					{Suit: Diamonds, Rank: Seven},
					{Suit: Clubs, Rank: Seven},
				},
			},
			true,
		},
		{
			"Happy test",
			fields{nil},
			args{
				topCards: []BigTwoCard{
					{Suit: Diamonds, Rank: Five},
					{Suit: Hearts, Rank: Five},
					{Suit: Spades, Rank: Five},
					{Suit: Diamonds, Rank: Two},
					{Suit: Clubs, Rank: Two},
				},
				playCards: []BigTwoCard{
					{Suit: Spades, Rank: Ace},
					{Suit: Diamonds, Rank: Ace},
					{Suit: Hearts, Rank: Ace},
					{Suit: Spades, Rank: Three},
					{Suit: Hearts, Rank: Three},
				},
			},
			true,
		},
		{
			"Unhappy test",
			fields{nil},
			args{
				topCards: []BigTwoCard{
					{Suit: Spades, Rank: Two},
					{Suit: Diamonds, Rank: Two},
					{Suit: Hearts, Rank: Two},
					{Suit: Spades, Rank: Three},
					{Suit: Hearts, Rank: Three},
				},
				playCards: []BigTwoCard{
					{Suit: Diamonds, Rank: Three},
					{Suit: Hearts, Rank: Three},
					{Suit: Spades, Rank: Three},
					{Suit: Diamonds, Rank: Four},
					{Suit: Clubs, Rank: Four},
				},
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
