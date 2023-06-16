package service

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnoGame_ShuffleDeck(t *testing.T) {
	deck := NewDeck()

	game := NewUnoGame(nil, deck)
	game.ShuffleDeck()

	// Since shuffling the deck is a random process, we can only assert that the number
	// of cards remains the same after shuffling.
	assert.Len(t, game.Deck.Cards, len(deck.Cards))

	// Assert that the deck cards have been shuffled
	assert.NotEqual(t, game.Deck.Cards, NewDeck().Cards)
}

func TestUnoGame_DealHands(t *testing.T) {
	deck := NewDeck()

	players := []Player{
		&HumanPlayer{Name: "Player 1"},
		&ComputerPlayer{Name: "Computer 1"},
	}

	game := NewUnoGame(players, deck)
	game.DealHands(2)

	// Each player should have received 2 cards.
	assert.Len(t, players[0].GetHand(), 2)
	assert.Len(t, players[1].GetHand(), 2)
}

func TestUnoGame_isValidMove(t *testing.T) {
	// TODO: Add tests for the isValidMove method.
	t.Skip("Not implemented yet")
}

// TODO: Add more tests for the UnoGame class.
