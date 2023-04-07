package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"showdown/entity"
)

func TestNewGame(t *testing.T) {
	p1 := entity.NewHumanPlayer(1)
	p2 := entity.NewHumanPlayer(2)
	p3 := entity.NewHumanPlayer(3)
	pAI := entity.NewAIPlayer(4)
	var deck *entity.Deck
	var game *Game

	t.Run("Test creating game with human player, AI player, and new deck", func(t *testing.T) {
		deck = entity.NewDeck()
		game = NewGame(p1, p2, p3, pAI, deck)

		assert.IsType(t, &Game{}, game)
		assert.Equal(t, 4, len(game.Players))

		assert.Equal(t, 52, len(deck.Cards))
		assert.Equal(t, entity.Card{Suit: entity.Spades, Rank: entity.Ace}, deck.Cards[0])
	})

	t.Run("should successfully rename the human player", func(t *testing.T) {
		p1.SetName("TestPlayer1")
		p2.SetName("TestPlayer2")
		p3.SetName("TestPlayer3")

		assert.Equal(t, "TestPlayer1", p1.Name())
		assert.Equal(t, "TestPlayer2", p2.Name())
		assert.Equal(t, "TestPlayer3", p3.Name())
		assert.Equal(t, "AI has no name", pAI.Name())
	})

	t.Run("cards in a shuffled deck should be random ordered", func(t *testing.T) {
		deck.Shuffle()

		assert.NotEqual(t, entity.Card{Suit: entity.Spades, Rank: entity.Ace}, deck.Cards[0])
	})
}
