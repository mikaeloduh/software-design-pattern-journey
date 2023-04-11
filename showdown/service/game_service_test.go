package service

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"showdown/entity"
)

func TestNewGame(t *testing.T) {
	p1 := entity.NewHumanPlayer(0)
	p2 := entity.NewHumanPlayer(1)
	p3 := entity.NewHumanPlayer(2)
	pAI := entity.NewAIPlayer(3)
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
		game.Init()

		assert.NotEqual(t, entity.Card{Suit: entity.Spades, Rank: entity.Ace}, deck.Cards[0])
	})

	t.Run("when draw is finished, every player should have 13 hand card", func(t *testing.T) {
		game.DrawLoop()

		assert.IsType(t, entity.Card{}, p1.HandCards[0])
		assert.Equal(t, rounds, len(p1.HandCards))
		assert.Equal(t, rounds, len(p2.HandCards))
		assert.Equal(t, rounds, len(p3.HandCards))
		assert.Equal(t, rounds, len(pAI.HandCards))
		assert.Equal(t, 52-rounds*4, len(game.deck.Cards))
	})

	t.Run("Testing game over: game should be end after 13th rounds", func(t *testing.T) {
		game.takeTurnLoop()

		assert.Equal(t, 0, len(p1.HandCards))
		assert.Equal(t, 0, len(p2.HandCards))
		assert.Equal(t, 0, len(p3.HandCards))
		assert.Equal(t, 0, len(pAI.HandCards))
	})

	t.Run("Testing game result: winner's name and points", func(t *testing.T) {
		winner := game.gameResult()

		fmt.Printf("winner: %+v\n", winner)

		assert.NotEmpty(t, winner)
		for i := range game.Players {
			p := game.Players[i]
			if p != winner {
				fmt.Printf("looser: %+v\n", p)

				assert.Greater(t, winner.Point(), p.Point())
			}
		}
	})
}
