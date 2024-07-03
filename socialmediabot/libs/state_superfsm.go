package libs

import (
	"reflect"
)

// SuperFSM is a finite state machine, U is the subject type
type SuperFSM[U any] struct {
	Subject         U
	States          []IState
	currentStateIdx int
	transitions     []Transition
}

func NewSuperFSM[U any](subject U, initState IState) *SuperFSM[U] {
	return &SuperFSM[U]{
		Subject:         subject,
		States:          []IState{initState},
		currentStateIdx: 0,
		transitions:     make([]Transition, 0),
	}
}

func (m *SuperFSM[U]) Enter() {
	//TODO implement me
	panic("implement me")
}

func (m *SuperFSM[U]) Exit() {
	//TODO implement me
	panic("implement me")
}

func (m *SuperFSM[U]) GetState() IState {
	return m.currentState().GetState()
}

func (m *SuperFSM[U]) SetState(state IState) {
	for i, s := range m.States {
		if reflect.TypeOf(s) == reflect.TypeOf(state) {
			m.currentStateIdx = i
			break
		}
	}
}

func (m *SuperFSM[U]) AddState(state IState) {
	m.States = append(m.States, state)
}

func (m *SuperFSM[U]) AddTransition(transition *Transition) {
	m.transitions = append(m.transitions, *transition)
}

func (m *SuperFSM[U]) Trigger(event IEvent) {
	for _, transition := range m.transitions {
		if reflect.TypeOf(transition.event) == reflect.TypeOf(event) && reflect.TypeOf(transition.from) == reflect.TypeOf(m.currentState()) && transition.guard(event) {
			m.SetState(transition.to)
			transition.action()
			break
		}
	}

	m.currentState().Trigger(event)
}

func (m *SuperFSM[U]) currentState() IState {
	return m.States[m.currentStateIdx]
}
