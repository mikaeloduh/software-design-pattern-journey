package service

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"showdown/entity"
)

func TestNewGame(t *testing.T) {
	t.Run("test new game with player", func(t *testing.T) {
		p1 := entity.NewPlayer("TestName1")
		p2 := entity.NewPlayer("TestName2")
		p3 := entity.NewPlayer("TestName3")
		p4 := entity.NewPlayer("TestName4")
		game := NewGame(p1, p2, p3, p4)

		assert.IsType(t, &Game{}, game)
		assert.Equal(t, 4, len(game.Players))
	})
}
