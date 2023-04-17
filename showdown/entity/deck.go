package entity

import (
	"fmt"
	"math/rand"
)

type Suit int

const (
	Clubs    Suit = 0
	Diamonds Suit = 1
	Hearts   Suit = 2
	Spades   Suit = 3
)

type Rank int

const (
	Ace   Rank = 14
	Two   Rank = 15
	Three Rank = 3
	Four  Rank = 4
	Five  Rank = 5
	Six   Rank = 6
	Seven Rank = 7
	Eight Rank = 8
	Nine  Rank = 9
	Ten   Rank = 10
	Jack  Rank = 11
	Queen Rank = 12
	King  Rank = 13
)

type Card struct {
	Suit Suit
	Rank Rank
}

func (c *Card) IsGreater(other Card) bool {
	if c.Rank > other.Rank {
		return true
	} else if c.Rank < other.Rank {
		return false
	} else {
		if c.Suit > other.Suit {
			return true
		} else {
			return false
		}
	}
}

func (c *Card) String() string {
	var suitName string
	switch c.Suit {
	case Clubs:
		suitName = "♠"
	case Diamonds:
		suitName = "♦"
	case Hearts:
		suitName = "♥"
	case Spades:
		suitName = "♣"
	}
	var rankName string
	switch c.Rank {
	case Ace:
		rankName = "A"
	case Two:
		rankName = "2"
	case Three:
		rankName = "3"
	case Four:
		rankName = "4"
	case Five:
		rankName = "5"
	case Six:
		rankName = "6"
	case Seven:
		rankName = "7"
	case Eight:
		rankName = "8"
	case Nine:
		rankName = "9"
	case Ten:
		rankName = "10"
	case Jack:
		rankName = "J"
	case Queen:
		rankName = "Q"
	case King:
		rankName = "K"
	}
	return fmt.Sprintf("%s %s", rankName, suitName)
}

type Deck struct {
	Cards []Card
}

func (d *Deck) DrawCard() Card {
	card := d.Cards[0]
	d.Cards = d.Cards[1:]
	return card
}

func (d *Deck) Shuffle() {
	for i := range d.Cards {
		j := rand.Intn(i + 1)
		d.Cards[i], d.Cards[j] = d.Cards[j], d.Cards[i]
	}
}

func NewDeck() *Deck {
	deck := Deck{}
	for _, suit := range []Suit{Spades, Hearts, Diamonds, Clubs} {
		for rank := Three; rank <= Two; rank++ {
			deck.Cards = append(deck.Cards, Card{Rank: rank, Suit: suit})
		}
	}
	return &deck
}
