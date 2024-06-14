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

type InterfaceFSM[T any] interface {
	PublicMethod()
	GetState() InterfaceFSM[T]
}

// UnimplementedTestState
type UnimplementedTestState struct{}

func (UnimplementedTestState) PublicMethod() {
	fmt.Println("to be implement")
}

type DefaultTestState struct {
	writer io.Writer
	UnimplementedTestState
	SuperState[any, InterfaceFSM[any]]
}

func (s *DefaultTestState) GetState() InterfaceFSM[any] {
	return s
}

func (s *DefaultTestState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "DefaultTestState public method called")
}

// AnotherTestState
type AnotherTestState struct {
	writer io.Writer
	UnimplementedTestState
	SuperState[any, InterfaceFSM[any]]
}

func (s *AnotherTestState) GetState() InterfaceFSM[any] {
	return s
}

func (s *AnotherTestState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "AnotherTestState public method called")
}

// Test states in FSM
func TestFSM(t *testing.T) {
	t.Run("test creating new FSM should have an initial state", func(t *testing.T) {
		expectedState := &DefaultTestState{}
		fsm := NewFiniteStateMachine[any, InterfaceFSM[any]](nil, expectedState)

		currentState := fsm.GetState()

		assert.IsType(t, expectedState, currentState)
	})

	t.Run("test when an event occur meets the guardian criteria, it should trigger the transition", func(t *testing.T) {
		fsm := NewFiniteStateMachine[any, InterfaceFSM[any]](nil, &DefaultTestState{})

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		transition := NewTransition[InterfaceFSM[any]](&DefaultTestState{}, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(event)

		currentState := fsm.GetState()
		assert.IsType(t, expectedState, currentState)
	})

	t.Run("test when an event occur does not meet the guardian criteria, it should not trigger the transition", func(t *testing.T) {
		initState := &DefaultTestState{}
		fsm := NewFiniteStateMachine[any, InterfaceFSM[any]](nil, initState)

		event := Event("test-event")
		guard := func() bool { return false }
		action := func() {}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		transition := NewTransition[InterfaceFSM[any]](initState, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(event)

		currentState := fsm.GetState()
		assert.IsType(t, initState, currentState)
	})

	t.Run("test when an event does not meet any transition, it should not change the state", func(t *testing.T) {
		initState := &DefaultTestState{}
		fsm := NewFiniteStateMachine[any, InterfaceFSM[any]](nil, initState)

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := &AnotherTestState{}
		fsm.AddState(expectedState)
		transition := NewTransition[InterfaceFSM[any]](initState, expectedState, event, guard, action)
		fsm.AddTransition(transition)

		fsm.Trigger(Event("another-event"))

		currentState := fsm.GetState()
		assert.IsType(t, initState, currentState)
	})

	t.Run("test subject public method behavior should variate depends on it's state", func(t *testing.T) {
		var writer bytes.Buffer

		initState := &DefaultTestState{writer: &writer}
		fsm := NewFiniteStateMachine[any, InterfaceFSM[any]](nil, initState)

		event := Event("test-event")
		guard := func() bool { return true }
		action := func() {}
		expectedState := &AnotherTestState{writer: &writer}
		fsm.AddState(expectedState)
		fsm.AddTransition(NewTransition[InterfaceFSM[any]](initState, expectedState, event, guard, action))
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
	UnimplementedTestState
	SuperState[any, InterfaceFSM[any]]
}

func (s *DefaultConversationState) GetState() InterfaceFSM[any] {
	return s
}

func (s *DefaultConversationState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, " DefaultConversationState public method called")
}

// InteractiveState
type InteractiveState struct {
	writer io.Writer
	UnimplementedTestState
	SuperState[any, InterfaceFSM[any]]
}

func (s *InteractiveState) GetState() InterfaceFSM[any] {
	return s
}

func (s *InteractiveState) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "InteractiveState public method called")
}

// NormalTestFSM
type NormalTestFSM struct {
	writer io.Writer
	UnimplementedTestState
	FiniteStateMachine[any, InterfaceFSM[any]]
}

func (s *NormalTestFSM) PublicMethod() {
	_, _ = fmt.Fprint(s.writer, "NormalTestFSM public method called")
}

// RootTestFSM
type RootTestFSM struct {
	writer io.Writer
	UnimplementedTestState
	FiniteStateMachine[any, InterfaceFSM[any]]
}

// RecordFSM
type RecordFSM struct {
	writer io.Writer
	UnimplementedTestState
	FiniteStateMachine[any, InterfaceFSM[any]]
}

// Test Composite FSM
func TestFSM_Composite(t *testing.T) {
	t.Run("test Normal State contains Default Conversation State and Interactive State", func(t *testing.T) {
		var writer bytes.Buffer

		defaultConversationState := &DefaultConversationState{writer: &writer}
		interactiveState := &InteractiveState{writer: &writer}

		normalStateFSM := &NormalTestFSM{
			FiniteStateMachine: *NewFiniteStateMachine[any, InterfaceFSM[any]](nil, defaultConversationState),
		}
		normalStateFSM.AddState(interactiveState)

		recordStateFSM := &RecordFSM{}
		rootFSM := RootTestFSM{
			FiniteStateMachine: *NewFiniteStateMachine[any, InterfaceFSM[any]](nil, normalStateFSM),
		}
		rootFSM.AddState(recordStateFSM)

		currentState := rootFSM.GetState()

		//assert.IsType(t, interactiveState, currentState)
		assert.Same(t, defaultConversationState, currentState)
	})
}

//func TestFSM_Composite_2(t *testing.T) {
//	t.Run("test2 Normal State contains Default Conversation State and Interactive State", func(t *testing.T) {
//		var writer bytes.Buffer
//
//		testDefaultConversationState := &TestDefaultConversationState{writer: &writer}
//
//		testNormalStateFSM := TestNormalStateFSM{}
//		testNormalStateFSM.GetState()
//
//		testInteractiveState := &TestInteractiveState{writer: &writer}
//		testNormalStateFSM.AddState(testInteractiveState)
//
//		rootFSM :=
//		recordStateFSM := NewFiniteStateMachine[any, IFSM](nil, nil)
//		rootFSM.AddState(recordStateFSM)
//
//		currentState := rootFSM.GetState()
//
//		//assert.IsType(t, testInteractiveState, currentState)
//		assert.Same(t, testDefaultConversationState, currentState)
//	})
//}

// test helper

// FakeStatefulSubject
type FakeStatefulSubject struct {
	fsm    *FiniteStateMachine[any, InterfaceFSM[any]]
	writer io.Writer
}

func (s *FakeStatefulSubject) PublicMethod() {
	s.fsm.GetState().PublicMethod()
}
