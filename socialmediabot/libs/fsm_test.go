package libs

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

// ITestState
type ITestState interface {
	PublicMethod()
}

// UnimplementedTestState
type UnimplementedTestState struct{}

func (UnimplementedTestState) PublicMethod() {
	fmt.Println("to be implement")
}

type DefaultTestState struct {
	writer io.Writer
	SuperState[any]
	UnimplementedTestState
}

func (s *DefaultTestState) GetState() IFSM {
	return s
}

func (s *DefaultTestState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "DefaultTestState public method called")
}

// AnotherTestState
type AnotherTestState struct {
	writer io.Writer
	SuperState[any]
	UnimplementedTestState
}

func (s *AnotherTestState) GetState() IFSM {
	return s
}

func (s *AnotherTestState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "AnotherTestState public method called")
}

// Test states in FSM
func TestFSM(t *testing.T) {
	t.Run("test creating new FSM should have an initial state", func(t *testing.T) {
		expectedState := &DefaultTestState{}
		fsm := NewFiniteStateMachine[any](nil, expectedState)

		currentState := fsm.GetState()

		assert.IsType(t, expectedState, currentState)
	})

	t.Run("test when an event occur meets the guardian criteria, it should trigger the transition", func(t *testing.T) {
		fsm := NewFiniteStateMachine[any](nil, &DefaultTestState{})

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		transition := NewTransition(&DefaultTestState{}, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(event)

		currentState := fsm.GetState()
		assert.IsType(t, expectedState, currentState)
	})

	t.Run("test when an event occur does not meet the guardian criteria, it should not trigger the transition", func(t *testing.T) {
		initState := &DefaultTestState{}
		fsm := NewFiniteStateMachine[any](nil, initState)

		event := Event("test-event")
		guard := func() bool { return false }
		action := func() {}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		transition := NewTransition(initState, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(event)

		currentState := fsm.GetState()
		assert.IsType(t, initState, currentState)
	})

	t.Run("test when an event does not meet any transition, it should not change the state", func(t *testing.T) {
		initState := &DefaultTestState{}
		fsm := NewFiniteStateMachine[any](nil, initState)

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		transition := NewTransition(initState, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(Event("another-event"))

		currentState := fsm.GetState()
		assert.IsType(t, initState, currentState)
	})

	t.Run("test subject public method behavior should variate depends on it's state", func(t *testing.T) {
		var writer bytes.Buffer

		initState := &DefaultTestState{writer: &writer}
		fsm := NewFiniteStateMachine[any](nil, initState)

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := &AnotherTestState{writer: &writer}
		fsm.AddState(expectedState)
		fsm.AddTransition(NewTransition(initState, expectedState, event, guard, action))
		statefulSubject := FakeStatefulSubject{fsm: fsm, writer: &writer}
		statefulSubject.PublicMethod()

		// assert when calling the subject public method, the corresponding state public method should be called
		// (in this case, the normal state public method should be called)
		assert.Equal(t, "DefaultTestState public method called", writer.String())

		// reset the writer
		writer.Reset()

		// (another state public method should be called)
		fsm.Trigger(event)
		statefulSubject.PublicMethod()
		assert.Equal(t, "AnotherTestState public method called", writer.String())
	})
}

// DefaultConversationState
type DefaultConversationState struct {
	writer io.Writer
	SuperState[any]
	UnimplementedTestState
}

func (s *DefaultConversationState) GetState() IFSM {
	return s
}

func (s *DefaultConversationState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, " DefaultConversationState public method called")
}

// InteractiveState
type InteractiveState struct {
	writer io.Writer
	SuperState[any]
	UnimplementedTestState
}

func (s *InteractiveState) GetState() IFSM {
	return s
}

func (s *InteractiveState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "InteractiveState public method called")
}

// NormalTestFSM
type NormalTestFSM struct {
	writer io.Writer
	FiniteStateMachine[any]
	UnimplementedTestState
}

func (s *NormalTestFSM) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "NormalTestFSM public method called")
}

// RootTestFSM
type RootTestFSM struct {
	writer io.Writer
	FiniteStateMachine[any]
	UnimplementedTestState
}

// RecordFSM
type RecordFSM struct {
	writer io.Writer
	FiniteStateMachine[any]
	UnimplementedTestState
}

// Test Composite FSM
func TestFSM_Composite(t *testing.T) {
	t.Run("test NormalTestFSM contains DefaultConversationState and InteractiveState", func(t *testing.T) {
		var writer bytes.Buffer

		defaultConversationState := &DefaultConversationState{writer: &writer}
		interactiveState := &InteractiveState{writer: &writer}

		normalStateFSM := &NormalTestFSM{
			FiniteStateMachine: *NewFiniteStateMachine[any](nil, defaultConversationState),
		}
		normalStateFSM.AddState(interactiveState)

		recordStateFSM := &RecordFSM{}
		rootFSM := RootTestFSM{
			FiniteStateMachine: *NewFiniteStateMachine[any](nil, normalStateFSM),
		}
		rootFSM.AddState(recordStateFSM)

		currentState := rootFSM.GetState()

		assert.Same(t, defaultConversationState, currentState)
	})
}

// Test helpers

// FakeStatefulSubject
type FakeStatefulSubject struct {
	fsm    *FiniteStateMachine[any]
	writer io.Writer
}

func (s *FakeStatefulSubject) PublicMethod() {
	s.fsm.GetState().(ITestState).PublicMethod()
}
