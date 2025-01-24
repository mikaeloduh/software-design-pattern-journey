package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBasicAttack(t *testing.T) {
	t.Parallel()

	t.Run("test Hero BasicAttack based on it's STR", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit2 := NewHero("p2")
		unit2HP := unit2.CurrentHP

		_ = unit1.selectSkill(0)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

		assert.Equal(t, unit2HP-unit1.STR, unit2.CurrentHP)
	})
}
