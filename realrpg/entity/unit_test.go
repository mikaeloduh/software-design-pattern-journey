package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHero(t *testing.T) {
	t.Run("test Hero BasicAttack based on it's STR", func(t *testing.T) {
		unit1 := NewHero()
		unit2 := NewHero()
		unit2HP := unit2.HP

		action := unit1.SelectSkill(0)
		action.SelectTarget([]IUnit{unit2})
		action.Do()

		assert.Equal(t, unit2HP-unit1.STR, unit2.HP)
	})

	t.Run("test Hero WaterBall attack", func(t *testing.T) {
		w := &WaterBall{Damage: 120}
		unit1 := NewHero()
		unit1.AddSkill(w)
		unit2 := NewHero()
		unit2HP := unit2.HP

		action := unit1.SelectSkill(1)
		action.SelectTarget([]IUnit{unit2})
		action.Do()

		assert.Equal(t, unit2HP-w.Damage, unit2.HP)
	})
}
