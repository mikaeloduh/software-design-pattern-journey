package entity

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWaterBall(t *testing.T) {
	t.Parallel()

	t.Run("test Hero WaterBall attack", func(t *testing.T) {
		unit1 := NewHero("p1")
		w := NewWaterBall(unit1)
		unit1.AddSkill(w)
		unit2 := NewHero("p2")
		unit2HP := unit2.CurrentHP

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

		assert.Equal(t, unit2HP-w.Damage, unit2.CurrentHP)
	})

	t.Run("test Hero WaterBall attack should have damage and cost MP", func(t *testing.T) {
		unit1 := NewHero("p1")
		w := NewWaterBall(unit1)
		unit1.AddSkill(w)
		unit2 := NewHero("p2")
		unit2HP := unit2.CurrentHP
		unit1MP := unit1.CurrentMP

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

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
}
