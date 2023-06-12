package service

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUnoGame(t *testing.T) {
	t.Run("Hello world", func(t *testing.T) {
		assert.Equal(t, 2, 1+1)
	})

	t.Run("Test New Uno Game no error", func(t *testing.T) {
		game := NewUnoGame()
		game.Run()

		assert.Equal(t, 40, len(game.Deck.Cards))
	})
}
