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

		assert.IsType(t, &Slime{}, (*troop1)[1])
		assert.Contains(t, writer.String(), "Slime is taking action")
	})

	t.Run("test when Slime die, summoner should receive 30 HP", func(t *testing.T) {
		var writer bytes.Buffer
		unit1 := NewHero("p1")
		summon := FakeNewSummon(unit1, &writer)
		unit1.AddSkill(summon)
		unit1.SetHp(200)
		troop1 := NewTroop([]IUnit{unit1})

		unit1.selectSkill(1)
		unit1.doSkill()

		slime := (*troop1)[1]
		slime.TakeDamage(100)

		assert.Equal(t, 200+30, unit1.GetHp())
	})

	t.Run("test when cursed unit dies, the caster regains HP equal to the unit's current MP", func(t *testing.T) {
		unit1 := NewHero("p1")
		curse := NewCurse(unit1)
		unit1.AddSkill(curse)
		unit1.SetHp(20)
		unit1HP := unit1.GetHp()
		unit2 := NewHero("p2")
		unit2MP := unit2.GetMp()

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

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

		troop1 := NewTroop([]IUnit{unit1})
		_ = NewTroop([]IUnit{unit2})

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()
		slime := (*troop1)[1]

		unit1.selectSkill(2)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

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

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unitQQ})
		unit1.doSkill()

		unit2.selectSkill(1)
		unit2.selectTarget([]IUnit{unitQQ})
		unit2.doSkill()

		unit3.selectSkill(1)
		unit3.selectTarget([]IUnit{unitQQ})
		unit3.doSkill()

		// when cursed unit die
		unitQQ.TakeDamage(9999)

		assert.Equal(t, unit1Hp+unitQQMp, unit1.GetHp())
		assert.Equal(t, unit2Hp+unitQQMp, unit2.GetHp())
		assert.Equal(t, unit3Hp+unitQQMp, unit3.GetHp())
	})

	onePunchHandler := Hp500Handler{
		PetrochemicalStateHandler{
			PoisonedStateHandler{
				CheerUpStateHandler{
					NormalStateHandler{nil},
				},
			},
		},
	}

	t.Run("test OnePunch case 1", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewOnePunch(unit1, onePunchHandler))

		unit2 := NewHero("p2")
		unit2Hp := unit2.GetHp()

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2Hp-300, unit2.GetHp())
	})

	t.Run("test OnePunch case 2", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewOnePunch(unit1, onePunchHandler))

		unit2 := NewHero("p2")
		unit2.SetHp(499)
		unit2.SetState(NewPetrochemicalState(unit2))
		unit2Hp := unit2.GetHp()

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2Hp-80, unit2.GetHp())
	})

	t.Run("test OnePunch case 3", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewOnePunch(unit1, onePunchHandler))

		unit2 := NewHero("p2")
		unit2.SetHp(499)
		unit2.SetState(NewCheerUpState(unit2))
		unit2Hp := unit2.GetHp()

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2Hp-100, unit2.GetHp())
		assert.IsType(t, &NormalState{}, unit2.GetState())
	})

	t.Run("test OnePunch case 4", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewOnePunch(unit1, onePunchHandler))

		unit2 := NewHero("p2")
		unit2.SetHp(499)
		unit2Hp := unit2.GetHp()

		unit1.selectSkill(1)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.Equal(t, unit2Hp-100, unit2.GetHp())
	})

}

func FakeNewSummon(unit IUnit, w io.Writer) *Summon {
	return &Summon{
		MPCost: 150,
		unit:   unit,
		Writer: w,
	}
}
