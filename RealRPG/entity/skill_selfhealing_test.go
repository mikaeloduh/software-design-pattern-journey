package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelfHealing(t *testing.T) {
	t.Parallel()

	t.Run("test SelfHealing should self-heal 150 HP", func(t *testing.T) {
		unit1 := NewHero("p1")
		selfHealing := NewSelfHealing(unit1)
		unit1.AddSkill(selfHealing)
		unit1.SetHp(500)
		unit1HP := unit1.CurrentHP

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit1)
		_ = unit1.doSkill(unit1)

		assert.Equal(t, unit1HP-selfHealing.Damage, unit1.CurrentHP)
	})
}
