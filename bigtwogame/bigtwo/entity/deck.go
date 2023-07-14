package entity

import "bigtwogame/template"

// NewBigTwoDeck contains BigTwoCard
func NewBigTwoDeck() *template.Deck[BigTwoCard] {
	deck := &template.Deck[BigTwoCard]{}
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for rank := Three; rank <= Two; rank++ {
			deck.Cards = append(deck.Cards, BigTwoCard{Rank: rank, Suit: suit})
		}
	}
	return deck
}
