package entity

import (
	"math/rand"
	"time"
)

type UnoCard struct {
	Color string
	Value string
}

type UnoDeck struct {
	Cards []UnoCard
}

func NewUnoDeck() UnoDeck {
	colors := []string{"Red", "Blue", "Green", "Yellow"}
	values := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

	deck := UnoDeck{}
	for _, color := range colors {
		for _, value := range values {
			card := UnoCard{Color: color, Value: value}
			deck.Cards = append(deck.Cards, card)
		}
	}

	return deck
}

func (d *UnoDeck) Shuffle() {
	rand.Seed(time.Now().UnixNano())

	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func (d *UnoDeck) DrawCard() UnoCard {
	card := d.Cards[len(d.Cards)-1]
	d.Cards = d.Cards[:len(d.Cards)-1]
	return card
}
