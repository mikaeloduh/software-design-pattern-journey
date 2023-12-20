package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPoisonState(t *testing.T) {
	t.Parallel()

	t.Run("test PoisonedState", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.SetState(NewPoisonedState(unit1))

		assert.IsType(t, &PoisonedState{}, unit1.State)
	})

}
