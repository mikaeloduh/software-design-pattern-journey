package entity

import "bigtwogame/template"

type ShowDownDeck struct {
	template.Deck[ShowDownCard]
}

// NewShowdownDeck contains ShowDownCard
func NewShowdownDeck() *ShowDownDeck {
	deck := &ShowDownDeck{}
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for rank := Three; rank <= Two; rank++ {
			deck.Cards = append(deck.Cards, ShowDownCard{Rank: rank, Suit: suit})
		}
	}
	return deck
}
