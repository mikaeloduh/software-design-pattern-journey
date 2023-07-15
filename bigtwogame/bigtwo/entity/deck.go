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

// Pattern Handler
// challenger, topCard --> [IF both valid single, DO compare] --> [IF both valid pair, DO compare]
// challenger, null    --> [IF topCard is null & challenger valid single, DO pass]
//                     --> [IF topCard is null $ challenger valid pair, DO pass]
// challenger, Club 3  --> [IF topCard is Club 3 & challenger is Club 3, DO pass]
//                     --> [IF topCard is Club 3 & challenger contains C3 and valid pair, DO pass]

// 1. contain Club 3
// 2. other
