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

	t.Run("test Summon a Slime to join Troop", func(t *testing.T) {

		//summon := &Summon{MPCost: 150}
		//unit1 := NewHero("p1")
		//unit1.AddSkill(summon)
		//
		//troop := Troop{
		//	Roles: []IUnit{unit1},
		//}
		//
		//unit1.selectSkill(1)
		//
		//assert.Len(t, 2, troop.Roles)
	})

	t.Run("test Hero SelfExplosion should attack all unit in field and kill himself", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit2 := NewHero("p2")
		unit3 := NewHero("p3")
		rpg := NewRPG([]IUnit{unit1, unit2, unit3}, nil)
		unit2HP := unit2.GetHp()
		unit3HP := unit3.GetHp()

		selfExplosion := NewSelfExplosion(unit1)
		unit1.AddSkill(selfExplosion)
		unit1.selectSkill(1)
		unit1.doSkill()

		selfExplosion.SelectTarget(rpg.units)
		selfExplosion.Do()

		assert.Equal(t, unit2HP-selfExplosion.Damage, unit2.HP)
		assert.Equal(t, unit3HP-selfExplosion.Damage, unit3.HP)
		assert.Equal(t, 0, unit1.GetHp())
	})

	t.Run("test Hero's SelfHealing can heal HP", func(t *testing.T) {
		selfHealing := &SelfHealing{MPCost: 50, Damage: 50}
		unit1 := NewHero("p1")
		unit1.AddSkill(selfHealing)
		unit1HP := unit1.HP

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit1})
		unit1.doSkill()

		assert.Equal(t, unit1HP+selfHealing.Damage, unit1.HP)
	})
}
