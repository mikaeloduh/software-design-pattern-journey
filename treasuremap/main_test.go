package main

import (
	"github.com/stretchr/testify/assert"
	"treasuremap/entity"

	"bytes"
	"testing"
)

func TestCharacterStatus(t *testing.T) {

	t.Run("test when character poisoned, HP should -15 for 3 rounds, then return to NormalState", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewPoisonedState(c))

		assert.Equal(t, 300, c.Hp)
		assert.IsType(t, &entity.PoisonedState{}, c.State)

		g.StartRound()
		assert.Equal(t, 300-15, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-15-15, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-15-15-15, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-15-15-15, c.Hp)
		assert.IsType(t, &entity.NormalState{}, c.State)
	})

	t.Run("test character in InvincibleState should invulnerable from damage", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewInvincibleState(c))
		c.TakeDamage(10)

		assert.IsType(t, &entity.InvincibleState{}, c.State)
		assert.Equal(t, 300, c.Hp)

		g.StartRound()
		c.TakeDamage(10)

		assert.Equal(t, 300, c.Hp)
		assert.IsType(t, &entity.InvincibleState{}, c.State)

		g.StartRound()
		c.TakeDamage(10)

		assert.Equal(t, 300-10, c.Hp)
		assert.IsType(t, &entity.NormalState{}, c.State)
	})

	t.Run("test character in HealingState should restore 30 Hp for 5 rounds", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.TakeDamage(200)
		c.SetState(entity.NewHealingState(c))

		g.StartRound()
		assert.Equal(t, 300-200+30, c.Hp)
		assert.IsType(t, &entity.HealingState{}, c.State)

		g.StartRound()
		assert.Equal(t, 300-200+30+30, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-200+30+30+30, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-200+30+30+30+30, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-200+30+30+30+30+30, c.Hp)

		g.StartRound()
		assert.Equal(t, 300-200+30+30+30+30+30, c.Hp)
		assert.IsType(t, &entity.NormalState{}, c.State)
	})

	t.Run("test while character's Hp is fully restored in HealingState, then should recover to NormalState immediately", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.TakeDamage(10)

		c.SetState(entity.NewHealingState(c))
		g.StartRound()

		assert.Equal(t, 300, c.Hp)
		assert.IsType(t, &entity.NormalState{}, c.State)
	})

	t.Run("test while in AcceleratedState, character can take 2 actions", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewAcceleratedState(c))
		g.StartRound()

		assert.IsType(t, &entity.AcceleratedState{}, c.State)
		assert.Equal(t, "take action\ntake action\n", writer.String())
	})

	t.Run("test AccelerateState will be interrupted by attack", func(t *testing.T) {
		var writer bytes.Buffer

		c := entity.NewCharacter(&writer)
		g := entity.NewAdventureGame(c)

		c.SetState(entity.NewAcceleratedState(c))

		assert.IsType(t, &entity.AcceleratedState{}, c.State)
		assert.Equal(t, 2, c.Speed)

		g.StartRound()
		c.TakeDamage(1)

		assert.IsType(t, &entity.NormalState{}, c.State)
		assert.Equal(t, 1, c.Speed)
	})
}
