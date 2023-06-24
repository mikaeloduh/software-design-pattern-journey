package template

import (
	"math/rand"
	"time"
)

type Deck[T ICard] struct {
	Cards []T
}

// Shuffle randomly shuffles the deck of cards.
func (d *Deck[T]) Shuffle() {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.Cards), func(i, j int) {
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	})
}

// DealCard deals a card from the deck.
func (d *Deck[T]) DealCard() T {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}
