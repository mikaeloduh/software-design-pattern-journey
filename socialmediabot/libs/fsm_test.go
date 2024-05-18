package libs

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFSM(t *testing.T) {
	t.Run("test creating new FSM should have an initial state", func(t *testing.T) {
		expectedState := &NormalTestState{}
		fsm := NewFiniteStateMachine[any, ITestState](nil, expectedState)

		currentState := fsm.GetState()

		assert.IsType(t, expectedState, currentState)
	})

	t.Run("test when an event occur meets the guardian criteria, it should trigger the transition", func(t *testing.T) {
		fsm := NewFiniteStateMachine[any, ITestState](nil, &NormalTestState{})

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		transition := NewTransition[ITestState](&NormalTestState{}, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(event)

		currentState := fsm.GetState()
		assert.IsType(t, expectedState, currentState)
	})

	t.Run("test when an event occur does not meet the guardian criteria, it should not trigger the transition", func(t *testing.T) {
		initState := &NormalTestState{}
		fsm := NewFiniteStateMachine[any, ITestState](nil, initState)

		event := Event("test-event")
		guard := func() bool { return false }
		action := func() {}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		transition := NewTransition[ITestState](initState, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(event)

		currentState := fsm.GetState()
		assert.IsType(t, initState, currentState)
	})

	t.Run("test when an event does not meet any transition, it should not change the state", func(t *testing.T) {
		initState := &NormalTestState{}
		fsm := NewFiniteStateMachine[any, ITestState](nil, initState)

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		transition := NewTransition[ITestState](initState, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(Event("another-event"))

		currentState := fsm.GetState()
		assert.IsType(t, initState, currentState)
	})

	t.Run("test subject public method behavior should variate depends on it's state", func(t *testing.T) {
		var writer bytes.Buffer

		initState := &NormalTestState{writer: &writer}
		fsm := NewFiniteStateMachine[any, ITestState](nil, initState)

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := &AnotherTestState{writer: &writer}
		fsm.AddState(expectedState)
		fsm.AddTransition(NewTransition[ITestState](initState, expectedState, event, guard, action))
		statefulSubject := FakeStatefulSubject{fsm: fsm, writer: &writer}
		statefulSubject.PublicMethod()

		// assert when calling the subject public method, the corresponding state public method should be called
		// (in this case, the normal state public method should be called)
		assert.Equal(t, "normal state public method called", writer.String())

		// reset the writer
		writer.Reset()

		// (another state public method should be called)
		fsm.Trigger(event)
		statefulSubject.PublicMethod()
		assert.Equal(t, "another state public method called", writer.String())
	})
}

// test helper
type ITestState interface {
	PublicMethod()
}

type NormalTestState struct {
	writer io.Writer
}

func (s *NormalTestState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "normal state public method called")
}

type AnotherTestState struct {
	writer io.Writer
}

func (s *AnotherTestState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "another state public method called")
}

type FakeStatefulSubject struct {
	fsm    *FiniteStateMachine[any, ITestState]
	writer io.Writer
}

func (s *FakeStatefulSubject) PublicMethod() {
	s.fsm.GetState().PublicMethod()
}
