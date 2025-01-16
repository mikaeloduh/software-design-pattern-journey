package libs

import (
	"reflect"
)

// SuperFSM is a finite state machine, U is the subject type
type SuperFSM[T IState[T]] struct {
	states          []T
	currentStateIdx int
	transitions     []Transition[T]
}

func NewSuperFSM[T IState[T]](initState T) SuperFSM[T] {
	return SuperFSM[T]{
		states:          []T{initState},
		currentStateIdx: 0,
		transitions:     make([]Transition[T], 0),
	}
}

func (m *SuperFSM[T]) Enter(_ IEvent) {
	m.Trigger(EnterStateEvent{})
}

func (m *SuperFSM[T]) Exit() {
	m.currentState().Exit()
}

func (m *SuperFSM[T]) GetState() T {
	return m.currentState().GetState()
}

func (m *SuperFSM[T]) SetState(state T, event IEvent) {
	for i, s := range m.states {
		if reflect.TypeOf(s) == reflect.TypeOf(state) {
			m.currentState().Exit()
			m.currentStateIdx = i
			m.currentState().Enter(event)
			break
		}
	}
}

func (m *SuperFSM[T]) AddState(state ...T) {
	m.states = append(m.states, state...)
}

func (m *SuperFSM[T]) AddTransition(transitions ...Transition[T]) {
	m.transitions = append(m.transitions, transitions...)
}

func (m *SuperFSM[T]) Trigger(event IEvent) {
	for _, transition := range m.transitions {
		if reflect.TypeOf(transition.event) == reflect.TypeOf(event) && reflect.TypeOf(transition.from) == reflect.TypeOf(m.currentState()) && transition.guard(event) {
			m.SetState(transition.to, event)
			transition.action(event)
			break
		}
	}

	m.currentState().Trigger(event)
}

func (m *SuperFSM[T]) currentState() T {
	return m.states[m.currentStateIdx]
}
