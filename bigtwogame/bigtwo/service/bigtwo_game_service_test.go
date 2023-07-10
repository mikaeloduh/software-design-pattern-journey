package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"bigtwogame/bigtwo/entity"
	"bigtwogame/template"
)

func TestBigTwo(t *testing.T) {
	t.Parallel()

	players := []entity.IBigTwoPlayer[entity.BigTwoCard]{
		&entity.AiBigTwoPlayer{Name: "Computer 1"},
		&entity.AiBigTwoPlayer{Name: "Computer 2"},
		&entity.AiBigTwoPlayer{Name: "Computer 3"},
		&entity.AiBigTwoPlayer{Name: "Computer 4"},
	}

	t.Run("New game success and have 4 players", func(t *testing.T) {
		game := NewBigTwoGame(players)

		assert.Equal(t, 4, len(game.Players))
	})

	t.Run("New a Deck and have it shuffled", func(t *testing.T) {
		deck := entity.NewBigTwoDeck()
		deck.Shuffle()

		assert.Equal(t, 52, len(deck.Cards))
		assert.NotEqual(t, entity.NewBigTwoDeck(), deck)
	})

	t.Run("New game and have card deal to all players", func(t *testing.T) {
		game := NewBigTwoGame(players)
		game.ShuffleDeck()
		game.DrawHands(game.NumCard)

		assert.Equal(t, 13, len(game.Players[0].GetHand()))
		assert.Equal(t, 13, len(game.Players[1].GetHand()))
		assert.Equal(t, 13, len(game.Players[2].GetHand()))
		assert.Equal(t, 13, len(game.Players[3].GetHand()))
	})

	t.Run("PreTakeTurn should play ♣3", func(t *testing.T) {
		deck := entity.NewBigTwoDeck()
		playingGame := &BigTwoGame{Players: players, Deck: deck}
		game := &template.GameFramework[entity.BigTwoCard]{
			Deck:        deck,
			Players:     make([]template.IPlayer[entity.BigTwoCard], len(players)),
			NumCard:     13,
			PlayingGame: playingGame,
		}
		for i, player := range players {
			game.Players[i] = player
		}

		assert.Equal(t, entity.BigTwoCard{}, playingGame.DeskCard)

		playingGame.PreTakeTurns()

		assert.Equal(t, entity.Clubs, playingGame.DeskCard.Suit)
		assert.Equal(t, entity.Three, playingGame.DeskCard.Rank)
	})
}
