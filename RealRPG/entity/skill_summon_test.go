package entity

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"io"
	"testing"
)

func TestSummon(t *testing.T) {
	t.Parallel()

	t.Run("test Summon a Slime to join Troop", func(t *testing.T) {
		unit1 := NewHero("p1")
		summon := NewSummon(unit1)
		unit1.AddSkill(summon)

		troop1 := NewTroop(unit1)
		assert.Equal(t, 1, troop1.Len())

		_ = unit1.selectSkill(1)
		_ = unit1.doSkill()

		assert.Equal(t, 2, troop1.Len())
	})

	t.Run("test summoned Slime should be able to take action in current round", func(t *testing.T) {
		var writer bytes.Buffer

		unit1 := NewHero("p1")
		summon := FakeNewSummon(unit1, &writer)
		unit1.AddSkill(summon)
		unit2 := NewHero("p2")

		troop1 := NewTroop(unit1)
		troop2 := NewTroop(unit2)

		battle := NewBattle(troop1, troop2)

		_ = unit1.selectSkill(1)
		_ = unit1.doSkill()

		battle.OnRoundStart()

		assert.IsType(t, &Slime{}, (*troop1)[1])
		assert.Contains(t, writer.String(), "Slime is taking action")
	})

	t.Run("test when Slime die, summoner should receive 30 HP", func(t *testing.T) {
		var writer bytes.Buffer
		unit1 := NewHero("p1")
		summon := FakeNewSummon(unit1, &writer)
		unit1.AddSkill(summon)
		unit1.SetHp(200)
		troop1 := NewTroop(unit1)

		_ = unit1.selectSkill(1)
		_ = unit1.doSkill()

		slime := (*troop1)[1]
		slime.TakeDamage(100)

		assert.Equal(t, 200+30, unit1.GetHp())
	})
}

func FakeNewSummon(unit IUnit, w io.Writer) *Summon {
	return &Summon{
		MPCost: 150,
		unit:   unit,
		Writer: w,
	}
}
