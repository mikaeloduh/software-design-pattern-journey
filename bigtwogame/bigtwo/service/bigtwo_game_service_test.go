package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"bigtwogame/bigtwo/entity"
)

func TestBigTwo(t *testing.T) {
	players := []entity.IBigTwoPlayer[entity.BigTwoCard]{
		&entity.AiBigTwoPlayer{Name: "Computer 1"},
		&entity.AiBigTwoPlayer{Name: "Computer 2"},
		&entity.AiBigTwoPlayer{Name: "Computer 3"},
		&entity.AiBigTwoPlayer{Name: "Computer 4"},
	}

	t.Run("New game success and have 4 players", func(t *testing.T) {
		game := NewBigTwoGame(players)

		assert.IsType(t, &BigTwoGame{}, game)
		assert.Equal(t, 4, len(game.Players))
	})

	t.Run("New a Deck and have it shuffled", func(t *testing.T) {
		deck := entity.NewBigTwoDeck()
		deck.Shuffle()

		assert.Equal(t, 52, len(deck.Cards))
		assert.NotEqual(t, entity.NewBigTwoDeck(), deck)
	})
}
