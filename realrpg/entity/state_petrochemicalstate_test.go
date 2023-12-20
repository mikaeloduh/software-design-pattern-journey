package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPetrochemicalState(t *testing.T) {
	t.Parallel()

	t.Run("test PetrochemicalState", func(t *testing.T) {
		unit1 := NewHero("p1")
		unit1.SetState(NewPetrochemicalState(unit1))

		assert.IsType(t, &PetrochemicalState{}, unit1.State)
	})

}
