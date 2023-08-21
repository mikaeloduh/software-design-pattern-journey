package main

import (
	"github.com/stretchr/testify/assert"
	"treasuremap/entity"

	"bytes"
	"testing"
)

func TestCharacterStatus(t *testing.T) {

	t.Run("test when character poisoned, HP should -15 for each following two round", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewPoisonedState(c))

		assert.Equal(t, "The character is poisoned.\n", writer.String())
		assert.Equal(t, 300, c.Hp)

		g.StartRound()

		assert.Equal(t, 300-15, c.Hp)

		g.StartRound()

		assert.Equal(t, 300-15-15, c.Hp)
	})
}
