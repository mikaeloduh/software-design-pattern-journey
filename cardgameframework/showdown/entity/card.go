package entity

import "fmt"

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
