package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnePunch(t *testing.T) {
	t.Parallel()

	onePunchHandler := Hp500Handler{
		PetrochemicalStateHandler{
			PoisonedStateHandler{
				CheerUpStateHandler{
					NormalStateHandler{nil}}}},
	}

	t.Run("test OnePunch case 1", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewOnePunch(unit1, onePunchHandler))

		unit2 := NewHero("p2")
		unit2Hp := unit2.GetHp()

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

		assert.Equal(t, unit2Hp-300, unit2.GetHp())
	})

	t.Run("test OnePunch case 2", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewOnePunch(unit1, onePunchHandler))

		unit2 := NewHero("p2")
		unit2.SetHp(499)
		unit2.SetState(NewPetrochemicalState(unit2))
		unit2Hp := unit2.GetHp()

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

		assert.Equal(t, unit2Hp-80, unit2.GetHp())
	})

	t.Run("test OnePunch case 3", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewOnePunch(unit1, onePunchHandler))

		unit2 := NewHero("p2")
		unit2.SetHp(499)
		unit2.SetState(NewCheerUpState(unit2))
		unit2Hp := unit2.GetHp()

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

		assert.Equal(t, unit2Hp-100, unit2.GetHp())
		assert.IsType(t, &NormalState{}, unit2.GetState())
	})

	t.Run("test OnePunch case 4", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewOnePunch(unit1, onePunchHandler))

		unit2 := NewHero("p2")
		unit2.SetHp(499)
		unit2Hp := unit2.GetHp()

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

		assert.Equal(t, unit2Hp-100, unit2.GetHp())
	})
}
