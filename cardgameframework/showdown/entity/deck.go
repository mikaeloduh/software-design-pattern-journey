package entity

import (
	"fmt"
	"math/rand"
)

// Suit type
type Suit int

const (
	Clubs    Suit = 0
	Diamonds Suit = 1
	Hearts   Suit = 2
	Spades   Suit = 3
)

func (s Suit) String() string {
	return map[Suit]string{
		Clubs:    "♠",
		Diamonds: "♦",
		Hearts:   "♥",
		Spades:   "♣",
	}[s]
}

// Rank type
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

func (r Rank) String() string {
	return map[Rank]string{
		Ace:   "A",
		Two:   "2",
		Three: "3",
		Four:  "4",
		Five:  "5",
		Six:   "6",
		Seven: "7",
		Eight: "8",
		Nine:  "9",
		Ten:   "10",
		Jack:  "J",
		Queen: "Q",
		King:  "K",
	}[r]
}

// Card type
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
		return c.Suit > other.Suit
	}
}

func (c *Card) String() string {
	return fmt.Sprintf("%s %s", c.Rank.String(), c.Suit.String())
}

// Deck contains Cards
type Deck struct {
	Cards []Card
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

func (d *Deck) DealCard() Card {
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
