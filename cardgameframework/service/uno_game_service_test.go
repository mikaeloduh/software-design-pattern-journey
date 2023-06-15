package service

import (
	"cardgameframework/entity"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUnoGame(t *testing.T) {
	t.Run("Hello world", func(t *testing.T) {
		assert.Equal(t, 2, 1+1)
	})

	t.Run("Test New Uno Game no error", func(t *testing.T) {
		p1 := entity.NewUnoAIPlayer()
		p2 := entity.NewUnoAIPlayer()
		p3 := entity.NewUnoAIPlayer()
		p4 := entity.NewUnoAIPlayer()
		game := NewUnoGame(p1, p2, p3, p4)
		game.Init()

		assert.Equal(t, 40, len(game.Deck.Cards))
	})

	t.Run("Test Draw", func(t *testing.T) {
		p1 := entity.NewUnoAIPlayer()
		p2 := entity.NewUnoAIPlayer()
		p3 := entity.NewUnoAIPlayer()
		p4 := entity.NewUnoAIPlayer()
		game := NewUnoGame(p1, p2, p3, p4)
		game.Draw()

		assert.Equal(t, 40, len(game.Deck.Cards))
	})
}
