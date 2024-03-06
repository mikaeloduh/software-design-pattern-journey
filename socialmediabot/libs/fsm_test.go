package libs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFSM(t *testing.T) {
	t.Run("test creating new FSM should have an initial state", func(t *testing.T) {
		fsm := NewFiniteStateMachine(NormalState{})

		currentState := fsm.GetState()

		assert.IsType(t, NormalState{}, currentState)
	})

	t.Run("test when an event occur meets the guardian criteria, it should trigger the transition", func(t *testing.T) {
		fsm := NewFiniteStateMachine(NormalState{})

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := AnotherState{}
		transition := NewTransition(NormalState{}, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(event)

		currentState := fsm.GetState()
		assert.IsType(t, expectedState, currentState)
	})
}

// test helper
type NormalState struct {
}

type AnotherState struct {
}
