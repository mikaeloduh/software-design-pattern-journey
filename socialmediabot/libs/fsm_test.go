package libs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFSM(t *testing.T) {
	t.Run("test fsm", func(t *testing.T) {
		fsm := NewFiniteStateMachine(NormalState{})

		currentState := fsm.GetState()

		assert.IsType(t, NormalState{}, currentState)
	})
}
