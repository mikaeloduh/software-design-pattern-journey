package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSelfExplosion(t *testing.T) {
	t.Parallel()

	t.Run("test Hero SelfExplosion should attack all unit in field and kill himself", func(t *testing.T) {
		unit1 := NewHero("the boomer")
		unit2 := NewHero("target 1")
		unit3 := NewHero("target 2")
		unit2HP := unit2.GetHp()
		unit3HP := unit3.GetHp()

		selfExplosion := NewSelfExplosion(unit1)
		unit1.AddSkill(selfExplosion)
		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2, unit3)
		_ = unit1.doSkill(unit2, unit3)

		assert.Equal(t, unit2HP-selfExplosion.Damage, unit2.GetHp())
		assert.Equal(t, unit3HP-selfExplosion.Damage, unit3.CurrentHP)
		assert.Equal(t, 0, unit1.GetHp())
	})
}
