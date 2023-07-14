package entity

import "bigtwogame/template"

// NewShowdownDeck contains ShowDownCard
func NewShowdownDeck() *template.Deck[ShowDownCard] {
	deck := &template.Deck[ShowDownCard]{}
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for rank := Three; rank <= Two; rank++ {
			deck.Cards = append(deck.Cards, ShowDownCard{Rank: rank, Suit: suit})
		}
	}
	return deck
}
