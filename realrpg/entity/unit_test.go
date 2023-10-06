package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHero(t *testing.T) {
	t.Run("test Hero BasicAttack based on it's STR", func(t *testing.T) {
		unit1 := NewHero()
		unit2 := NewHero()

		action := unit1.SelectSkill(0)
		action.SelectTarget([]IUnit{unit2})
		action.Do()

		assert.Equal(t, 1000-50, unit2.HP)
	})
}
