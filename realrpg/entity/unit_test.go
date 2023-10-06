package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHero(t *testing.T) {
	t.Run("test Hero BasicAttack based on it's STR", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit2 := NewHero("p2")
		unit2HP := unit2.HP

		unit1.selectSkill(0)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2HP-unit1.STR, unit2.HP)
	})

	t.Run("test Hero WaterBall attack", func(t *testing.T) {
		w := &WaterBall{Damage: 120}
		unit1 := NewHero("p1")
		unit1.AddSkill(w)
		unit2 := NewHero("p2")
		unit2HP := unit2.HP

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2HP-w.Damage, unit2.HP)
	})

	t.Run("test Hero WaterBall attack should have damage and cost MP", func(t *testing.T) {
		w := &WaterBall{Damage: 120, MPCost: 50}
		unit1 := NewHero("p1")
		unit1.AddSkill(w)
		unit2 := NewHero("p2")
		unit2HP := unit2.HP
		unit1MP := unit1.MP

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2HP-w.Damage, unit2.HP)
		assert.Equal(t, unit1MP-w.MPCost, unit1.MP)
	})

	t.Run("test Hero should have enough MP to select WaterBall", func(t *testing.T) {
		w := &WaterBall{Damage: 120, MPCost: 50}
		unit1 := NewHero("p1")
		unit1.AddSkill(w)
		unit1.MP = 30

		_, err := unit1.selectSkill(1)

		assert.Error(t, err)
	})
}
