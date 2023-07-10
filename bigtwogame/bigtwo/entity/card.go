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
	return fmt.Sprintf("%d", r)
}

// BigTwoCard type
type BigTwoCard struct {
	Suit Suit
	Rank Rank
}

func (c BigTwoCard) String() string {
	return fmt.Sprintf("%s %s", c.Rank.String(), c.Suit.String())
}

func (c BigTwoCard) IsGreater(other BigTwoCard) bool {
	if c.Rank > other.Rank {
		return true
	} else if c.Rank < other.Rank {
		return false
	} else {
		return c.Suit > other.Suit
	}
}
