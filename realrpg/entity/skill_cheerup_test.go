package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheerUp(t *testing.T) {
	t.Parallel()

	t.Run("test CheerUp targets should be in CheerUpState", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.AddSkill(NewCheerUp(unit1))
		unit2 := NewHero("p2")

		_ = unit1.selectSkill(1)
		_ = unit1.selectTarget(unit2)
		_ = unit1.doSkill(unit2)

		assert.IsType(t, &CheerUpState{}, unit2.GetState())
	})
}
