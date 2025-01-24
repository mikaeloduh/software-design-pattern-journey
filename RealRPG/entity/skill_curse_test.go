package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCurse(t *testing.T) {
	t.Parallel()

	t.Run("test when cursed unit dies, the caster regains HP equal to the unit's current MP", func(t *testing.T) {
		unit1 := NewHero("p1")
		curse := NewCurse(unit1)
		unit1.AddSkill(curse)
		unit1.SetHp(20)
		unit1HP := unit1.GetHp()
		unit2 := NewHero("p2")
		unit2MP := unit2.GetMp()

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

		unit2.TakeDamage(1000)

		assert.Equal(t, unit1HP+unit2MP, unit1.GetHp())
	})

	t.Run("test Summon + Curse together: when Slime and cursed Unit die, Hero should receive correct among of HP", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewSummon(unit1))
		unit1.AddSkill(NewCurse(unit1))
		unit1.SetHp(20)
		unit1HP := unit1.GetHp()

		unit2 := NewHero("p2")
		unit2MP := unit2.GetMp()

		troop1 := NewTroop(unit1)
		_ = NewTroop(unit2)

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)
		slime := (*troop1)[1]

		_ = unit1.selectSkill(2)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

		// when slime die
		slime.TakeDamage(9999)
		assert.Equal(t, unit1HP+30, unit1.GetHp())

		// when cursed unit die
		unit2.TakeDamage(9999)
		assert.Equal(t, unit1HP+30+unit2MP, unit1.GetHp())
	})

	t.Run("test a Unit should be able to cursed by multiple caster", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewCurse(unit1))
		unit2 := NewHero("p2")
		unit2.AddSkill(NewCurse(unit2))
		unit3 := NewHero("p3")
		unit3.AddSkill(NewCurse(unit3))
		unit1.SetHp(10)
		unit2.SetHp(10)
		unit3.SetHp(10)
		unit1Hp := unit1.GetHp()
		unit2Hp := unit2.GetHp()
		unit3Hp := unit3.GetHp()

		unitQQ := NewHero("a-poor-guy")
		unitQQMp := unitQQ.GetMp()

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unitQQ)
		_ = unit1.doSkill(unitQQ)

		_ = unit2.selectSkill(1)
		_ = unit2.selectTarget(unitQQ)
		_ = unit2.doSkill(unitQQ)

		_ = unit3.selectSkill(1)
		_ = unit3.selectTarget(unitQQ)
		_ = unit3.doSkill(unitQQ)

		// when cursed unit die
		unitQQ.TakeDamage(9999)

		assert.Equal(t, unit1Hp+unitQQMp, unit1.GetHp())
		assert.Equal(t, unit2Hp+unitQQMp, unit2.GetHp())
		assert.Equal(t, unit3Hp+unitQQMp, unit3.GetHp())
	})
}
