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

// TestStateFSM
type TestStateFSM struct {
	SuperFSM[any]
	UnimplementedTestState
}

func (f *TestStateFSM) PublicMethod() {
	f.GetState().(ITestState).PublicMethod()
}

// DefaultTestState
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

// TestEvent
type TestEvent struct {
}

func (t TestEvent) GetData() IEvent {
	panic("implement me")
}

// AnotherEvent
type AnotherEvent struct{}

func (AnotherEvent) GetData() IEvent {
	panic("implement me")
}

// TestPositiveGuard
type TestPositiveGuard struct {
}

func (t TestPositiveGuard) Exec(_ IEvent) bool {
	return true
}

// TestNegativeGuard
type TestNegativeGuard struct {
}

func (t TestNegativeGuard) Exec(_ IEvent) bool {
	return false
}

// Test states in FSM
func TestFSM(t *testing.T) {
	t.Run("test creating new FSM should have an initial state", func(t *testing.T) {
		expectedState := &DefaultTestState{}
		fsm := NewSuperFSM[any](nil, expectedState)

		assert.IsType(t, expectedState, fsm.GetState())
	})

	t.Run("test when an event occur meets the guardian criteria, it should trigger the transition", func(t *testing.T) {
		fsm := TestStateFSM{SuperFSM: *NewSuperFSM[any](nil, &DefaultTestState{})}

		testEvent := TestEvent{}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		fsm.AddTransition(NewTransition(&DefaultTestState{}, expectedState, testEvent, TestPositiveGuard{}, Action(func() {})))

		fsm.Trigger(testEvent)

		assert.IsType(t, expectedState, fsm.GetState())
	})

	t.Run("test when an event occur does not meet the guardian criteria, it should not trigger the transition", func(t *testing.T) {
		initState := &DefaultTestState{}
		fsm := TestStateFSM{SuperFSM: *NewSuperFSM[any](nil, initState)}

		testEvent := TestEvent{}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		fsm.AddTransition(NewTransition(initState, expectedState, testEvent, TestNegativeGuard{}, Action(func() {})))

		fsm.Trigger(testEvent)

		assert.IsType(t, initState, fsm.GetState())
	})

	t.Run("test when an event does not meet any transition, it should not change the state", func(t *testing.T) {
		initState := &DefaultTestState{}
		fsm := TestStateFSM{SuperFSM: *NewSuperFSM[any](nil, initState)}

		testEvent := TestEvent{}
		anotherState := &AnotherTestState{}
		fsm.AddState(anotherState)
		fsm.AddTransition(NewTransition(initState, anotherState, testEvent, TestPositiveGuard{}, Action(func() {})))

		anotherEvent := AnotherEvent{}
		fsm.Trigger(anotherEvent)

		assert.IsType(t, initState, fsm.GetState())
	})

	t.Run("test subject public method behavior should variate depends on it's state", func(t *testing.T) {
		var writer bytes.Buffer

		initState := &DefaultTestState{writer: &writer}
		fsm := TestStateFSM{SuperFSM: *NewSuperFSM[any](nil, initState)}

		testEvent := TestEvent{}
		expectedState := &AnotherTestState{writer: &writer}
		fsm.AddState(expectedState)
		fsm.AddTransition(NewTransition(initState, expectedState, testEvent, TestPositiveGuard{}, Action(func() {})))
		statefulSubject := FakeStatefulSubject{fsm: &fsm, writer: &writer}
		statefulSubject.PublicMethod()

		// assert when calling the subject public method, the corresponding state public method should be called
		// (in this case, the normal state public method should be called)
		assert.Equal(t, "DefaultTestState public method called", writer.String())
		writer.Reset()

		// (another state public method should be called)
		fsm.Trigger(testEvent)
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

	normalStateFSM := &NormalTestFSM{SuperFSM: *NewSuperFSM[any](nil, defaultConversationState)}
	normalStateFSM.AddState(interactiveState)

	recordStateFSM := &RecordFSM{}
	rootFSM := RootTestFSM{SuperFSM: *NewSuperFSM[any](nil, normalStateFSM)}
	rootFSM.AddState(recordStateFSM)

	t.Run("test NormalTestFSM contains DefaultConversationState and InteractiveState", func(t *testing.T) {
		currentState := rootFSM.GetState()

		assert.Same(t, defaultConversationState, currentState)
	})

	t.Run("calling public method from FSM is equivalent to calling the same method of the current state", func(t *testing.T) {
		normalStateFSM.PublicMethod()

		assert.Equal(t, "DefaultConversationState public method called", writer.String())
		writer.Reset()
	})

	t.Run("the behavior of public method should variant depends on it's current state", func(t *testing.T) {

		normalStateFSM.SetState(interactiveState)
		normalStateFSM.PublicMethod()

		assert.Equal(t, "InteractiveState public method called", writer.String())
	})
}

// Test helpers

// FakeStatefulSubject
type FakeStatefulSubject struct {
	fsm    *TestStateFSM
	writer io.Writer
}

func (s *FakeStatefulSubject) PublicMethod() {
	s.fsm.PublicMethod()
}
