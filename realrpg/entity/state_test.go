package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHero_State(t *testing.T) {
	t.Run("test Hero initial state is NormalState", func(t *testing.T) {
		unit1 := NewHero("p1")

		assert.IsType(t, &NormalState{}, unit1.State)
	})

	t.Run("test CheerUpState should have additional 50 damage on attack", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit2 := NewHero("p2")
		unit2HP := unit2.CurrentHP
		unit1.SetState(NewCheerUpState(unit1))

		unit1.selectSkill(0)
		unit1.selectTarget([]IUnit{unit2})
		unit1.doSkill()

		assert.IsType(t, &CheerUpState{}, unit1.State)
		assert.Equal(t, unit2HP-unit1.STR-50, unit2.CurrentHP)
	})
}
