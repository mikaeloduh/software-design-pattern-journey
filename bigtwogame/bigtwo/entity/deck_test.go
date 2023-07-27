package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSinglePattenComparator(t *testing.T) {
	type fields struct {
		Next PattenHandler
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
			h := SinglePattenComparator{
				Next: tt.fields.Next,
			}
			got := h.Do(tt.args.topCards, tt.args.playCards)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestSinglePattenValidator(t *testing.T) {
	type fields struct {
		Next PattenValidator
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
			v := SinglePattenValidator{
				Next: tt.fields.Next,
			}
			got := v.Do(tt.args.cards)
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestStraightPattenValidator(t *testing.T) {
	type fields struct {
		Next PattenValidator
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
			v := StraightPattenValidator{Next: tt.fields.Next}
			assert.Equalf(t, tt.want, v.Do(tt.args.cards), "Do(%v)", tt.args.cards)
		})
	}
}
