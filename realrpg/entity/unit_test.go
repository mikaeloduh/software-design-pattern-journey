package entity

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHero_skill(t *testing.T) {
	t.Run("test Hero BasicAttack based on it's STR", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit2 := NewHero("p2")
		unit2HP := unit2.CurrentHP

		unit1.selectSkill(0)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2HP-unit1.STR, unit2.CurrentHP)
	})

	t.Run("test Hero WaterBall attack", func(t *testing.T) {
		unit1 := NewHero("p1")
		w := NewWaterBall(unit1)
		unit1.AddSkill(w)
		unit2 := NewHero("p2")
		unit2HP := unit2.CurrentHP

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2HP-w.Damage, unit2.CurrentHP)
	})

	t.Run("test Hero WaterBall attack should have damage and cost MP", func(t *testing.T) {
		unit1 := NewHero("p1")
		w := NewWaterBall(unit1)
		unit1.AddSkill(w)
		unit2 := NewHero("p2")
		unit2HP := unit2.CurrentHP
		unit1MP := unit1.CurrentMP

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2HP-w.Damage, unit2.CurrentHP)
		assert.Equal(t, unit1MP-w.MPCost, unit1.CurrentMP)
	})

	t.Run("test Hero should have enough CurrentMP to select WaterBall", func(t *testing.T) {
		unit1 := NewHero("p1")
		w := NewWaterBall(unit1)
		unit1.AddSkill(w)
		unit1.CurrentMP = 30

		err := unit1.selectSkill(1)

		if assert.Error(t, err) {
			assert.Equal(t, err, fmt.Errorf("not enough CurrentMP"))
		}
	})

	t.Run("test Hero SelfExplosion should attack all unit in field and kill himself", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit2 := NewHero("p2")
		unit3 := NewHero("p3")
		troop1 := Troop{unit1}
		troop2 := Troop{unit2, unit3}
		_ = NewBattle(&troop1, &troop2)
		unit2HP := unit2.GetHp()
		unit3HP := unit3.GetHp()

		selfExplosion := NewSelfExplosion(unit1)
		unit1.AddSkill(selfExplosion)
		unit1.selectSkill(1)
		unit1.doSkill()

		selfExplosion.SelectTarget(troop2)
		selfExplosion.Do()

		assert.Equal(t, unit2HP-selfExplosion.Damage, unit2.CurrentHP)
		assert.Equal(t, unit3HP-selfExplosion.Damage, unit3.CurrentHP)
		assert.Equal(t, 0, unit1.GetHp())
	})

	t.Run("test SelfHealing should self-heal 150 HP", func(t *testing.T) {
		unit1 := NewHero("p1")
		selfHealing := NewSelfHealing(unit1)
		unit1.AddSkill(selfHealing)
		unit1.SetHp(500)
		unit1HP := unit1.CurrentHP

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit1})
		unit1.doSkill()

		assert.Equal(t, unit1HP-selfHealing.Damage, unit1.CurrentHP)
	})

	t.Run("test CheerUp targets should be in CheerUpState", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewCheerUp(unit1))
		unit2 := NewHero("p2")

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.IsType(t, &CheerUpState{}, unit2.GetState())
	})

	t.Run("test Summon a Slime to join Troop", func(t *testing.T) {
		unit1 := NewHero("p1")
		summon := NewSummon(unit1)
		unit1.AddSkill(summon)

		troop1 := NewTroop([]IUnit{unit1})
		assert.Equal(t, 1, troop1.Len())

		unit1.selectSkill(1)
		unit1.doSkill()

		assert.Equal(t, 2, troop1.Len())
	})

	t.Run("test summoned Slime should be able to take action in current round", func(t *testing.T) {
		var writer bytes.Buffer

		unit1 := NewHero("p1")
		summon := FakeNewSummon(unit1, &writer)
		unit1.AddSkill(summon)
		unit2 := NewHero("p2")

		troop1 := NewTroop([]IUnit{unit1})
		troop2 := NewTroop([]IUnit{unit2})

		battle := NewBattle(troop1, troop2)

		unit1.selectSkill(1)
		unit1.doSkill()

		battle.StartRound()

		assert.Contains(t, writer.String(), "Slime is taking action")
	})
}

func FakeNewSummon(unit IUnit, w io.Writer) *Summon {
	return &Summon{
		MPCost: 150,
		unit:   unit,
		Writer: w,
	}
}
