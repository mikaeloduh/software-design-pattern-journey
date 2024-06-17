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

func (s *DefaultTestState) GetState() IState {
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

func (s *AnotherTestState) GetState() IState {
	return s
}

func (s *AnotherTestState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "AnotherTestState public method called")
}

// Test states in FSM
func TestFSM(t *testing.T) {
	t.Run("test creating new FSM should have an initial state", func(t *testing.T) {
		expectedState := &DefaultTestState{}
		fsm := NewSuperFSM[any](nil, expectedState)

		currentState := fsm.GetState()

		assert.IsType(t, expectedState, currentState)
	})

	t.Run("test when an event occur meets the guardian criteria, it should trigger the transition", func(t *testing.T) {
		fsm := NewSuperFSM[any](nil, &DefaultTestState{})

		event := Event("test-event")
		guard := Guard(func() bool { return true })
		action := Action(func() {})
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
		fsm := NewSuperFSM[any](nil, initState)

		event := Event("test-event")
		guard := Guard(func() bool { return false })
		action := Action(func() {})
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
		fsm := NewSuperFSM[any](nil, initState)

		event := Event("test-event")
		guard := Guard(func() bool { return true })
		action := Action(func() {})
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
		fsm := NewSuperFSM[any](nil, initState)

		event := Event("test-event")
		guard := Guard(func() bool { return true })
		action := Action(func() {})
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

func (s *DefaultConversationState) GetState() IState {
	return s
}

func (s *DefaultConversationState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "DefaultConversationState public method called")
}

// InteractiveState
type InteractiveState struct {
	writer io.Writer
	SuperState[any]
	UnimplementedTestState
}

func (s *InteractiveState) GetState() IState {
	return s
}

func (s *InteractiveState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "InteractiveState public method called")
}

// NormalTestFSM
type NormalTestFSM struct {
	SuperFSM[any]
	UnimplementedTestState
}

func (s *NormalTestFSM) PublicMethod() {
	s.GetState().(ITestState).PublicMethod()
}

// RootTestFSM
type RootTestFSM struct {
	SuperFSM[any]
	UnimplementedTestState
}

// RecordFSM
type RecordFSM struct {
	SuperFSM[any]
	UnimplementedTestState
}

// Test Composite FSM
func TestFSM_Composite(t *testing.T) {
	var writer bytes.Buffer

	defaultConversationState := &DefaultConversationState{writer: &writer}
	interactiveState := &InteractiveState{writer: &writer}

	normalStateFSM := &NormalTestFSM{
		SuperFSM: *NewSuperFSM[any](nil, defaultConversationState),
	}
	normalStateFSM.AddState(interactiveState)

	recordStateFSM := &RecordFSM{}
	rootFSM := RootTestFSM{
		SuperFSM: *NewSuperFSM[any](nil, normalStateFSM),
	}
	rootFSM.AddState(recordStateFSM)

	t.Run("test NormalTestFSM contains DefaultConversationState and InteractiveState", func(t *testing.T) {

		currentState := rootFSM.GetState()

		assert.Same(t, defaultConversationState, currentState)
	})

	t.Run("calling public method from FSM is equivalent to calling the same method of the current state", func(t *testing.T) {
		normalStateFSM.PublicMethod()

		assert.Equal(t, "DefaultConversationState public method called", writer.String())
	})

	t.Run("the behavior of public method should variant depends on it's current state", func(t *testing.T) {
		writer.Reset()

		normalStateFSM.SetState(interactiveState)
		normalStateFSM.PublicMethod()

		assert.Equal(t, "InteractiveState public method called", writer.String())
	})
}

// Test helpers

// FakeStatefulSubject
type FakeStatefulSubject struct {
	fsm    *SuperFSM[any]
	writer io.Writer
}

func (s *FakeStatefulSubject) PublicMethod() {
	s.fsm.GetState().(ITestState).PublicMethod()
}
