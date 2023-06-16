package entity

import (
	"math/rand"
	"time"
)

// Card represents an UNO card.
type Card struct {
	Color string
	Value string
}

// Deck represents the UNO deck.
type Deck struct {
	Cards []Card
}

func NewDeck() Deck {
	colors := []string{"Red", "Blue", "Green", "Yellow"}
	values := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	deck := Deck{}
	for _, color := range colors {
		for _, value := range values {
			card := Card{Color: color, Value: value}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

// Shuffle randomly shuffles the deck of cards.
func (d *Deck) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// DealCard deals a card from the deck.
func (d *Deck) DealCard() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}
