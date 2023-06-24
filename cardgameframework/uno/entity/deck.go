package entity

import "cardgameframework/template"

// NewUnoDeck represents the UNO deck.
func NewUnoDeck() template.Deck[UnoCard] {
	deck := template.Deck[UnoCard]{}
	for _, color := range []Color{Red, Blue, Green, Yellow} {
		for value := Zero; value <= Nine; value++ {
			card := UnoCard{Color: color, Value: value}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}
