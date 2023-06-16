package service

import (
	"cardgameframework/uno/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnoGame_ShuffleDeck(t *testing.T) {
	deck := entity.NewDeck()

	game := NewUnoGame(nil, deck)
	game.ShuffleDeck()

	// Since shuffling the deck is a random process, we can only assert that the number
	// of cards remains the same after shuffling.
	assert.Len(t, game.Deck.Cards, len(deck.Cards))

	// Assert that the deck cards have been shuffled
	assert.NotEqual(t, game.Deck.Cards, entity.NewDeck().Cards)
}

func TestUnoGame_DealHands(t *testing.T) {
	deck := entity.NewDeck()

	players := []entity.IPlayer{
		&entity.HumanPlayer{Name: "IPlayer 1"},
		&entity.ComputerPlayer{Name: "Computer 1"},
	}

	game := NewUnoGame(players, deck)
	game.DealHands(2)

	// Each player should have received 2 cards.
	assert.Len(t, players[0].GetHand(), 2)
	assert.Len(t, players[1].GetHand(), 2)
}

func TestUnoGame_Result(t *testing.T) {
	deck := entity.NewDeck()
	players := []entity.IPlayer{
		&entity.ComputerPlayer{Name: "Computer 1"},
		&entity.ComputerPlayer{Name: "Computer 2"},
		&entity.ComputerPlayer{Name: "Computer 3"},
		&entity.ComputerPlayer{Name: "Computer 4"},
	}

	game := NewUnoGame(players, deck)
	game.ShuffleDeck()
	game.DealHands(5)
	game.TakeTurns()
	winner := game.GameResult()

	// IPlayer who won the game should have their hand empty
	assert.Equal(t, 0, len(winner.GetHand()))
}
