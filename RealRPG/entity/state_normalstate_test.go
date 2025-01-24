package entity

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNormalState(t *testing.T) {
	t.Parallel()

	t.Run("test Hero initial state is NormalState", func(t *testing.T) {
		unit1 := NewHero("p1")

		assert.IsType(t, &NormalState{}, unit1.State)
	})
}
