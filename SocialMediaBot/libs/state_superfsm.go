package libs

import (
	"reflect"
)

// SuperFSM is a finite state machine, U is the subject type
type SuperFSM[T IState[T]] struct {
	currentState T
	states       map[reflect.Type]T
	transitions  []Transition[T]
}

func NewSuperFSM[T IState[T]](initState T) SuperFSM[T] {
	fsm := SuperFSM[T]{
		currentState: initState,
		states:       map[reflect.Type]T{},
		transitions:  make([]Transition[T], 0),
	}
	fsm.states[reflect.TypeOf(initState)] = initState

	return fsm
}

func (m *SuperFSM[T]) Enter(_ IEvent) {
	m.Trigger(EnterStateEvent{})
}

func (m *SuperFSM[T]) Exit() {
	m.currentState.Exit()
}

func (m *SuperFSM[T]) GetState() T {
	return m.currentState.GetState()
}

func (m *SuperFSM[T]) SetState(state T, event IEvent) {
	if newState, ok := m.states[reflect.TypeOf(state)]; ok {
		m.currentState.Exit()
		m.currentState = newState
		m.currentState.Enter(event)
	}
}

func (m *SuperFSM[T]) AddState(state T) {
	m.states[reflect.TypeOf(state)] = state
}

func (m *SuperFSM[T]) AddTransition(from T, to T, event IEvent, guard Guard, action Action) {
	m.transitions = append(m.transitions, NewTransition(from, to, event, guard, action))
}

func (m *SuperFSM[T]) Trigger(event IEvent) {
	for _, transition := range m.transitions {
		if reflect.TypeOf(transition.event) == reflect.TypeOf(event) && reflect.TypeOf(transition.from) == reflect.TypeOf(m.currentState) && transition.guard(event) {
			m.SetState(transition.to, event)
			transition.action(event)
			break
		}
	}

	m.currentState.Trigger(event)
}
