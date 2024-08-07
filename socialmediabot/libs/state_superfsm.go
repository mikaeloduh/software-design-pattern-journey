package libs

import (
	"reflect"
)

// SuperFSM is a finite state machine, U is the subject type
type SuperFSM struct {
	states          []IState
	currentStateIdx int
	transitions     []Transition
}

func NewSuperFSM(initState IState) SuperFSM {
	return SuperFSM{
		states:          []IState{initState},
		currentStateIdx: 0,
		transitions:     make([]Transition, 0),
	}
}

func (m *SuperFSM) Enter(_ IEvent) {
	m.Trigger(EnterStateEvent{})
}

func (m *SuperFSM) Exit() {
	m.currentState().Exit()
}

func (m *SuperFSM) GetState() IState {
	return m.currentState().GetState()
}

func (m *SuperFSM) SetState(state IState, event IEvent) {
	for i, s := range m.states {
		if reflect.TypeOf(s) == reflect.TypeOf(state) {
			m.currentState().Exit()
			m.currentStateIdx = i
			m.currentState().Enter(event)
			break
		}
	}
}

func (m *SuperFSM) AddState(state ...IState) {
	m.states = append(m.states, state...)
}

func (m *SuperFSM) AddTransition(transitions ...Transition) {
	m.transitions = append(m.transitions, transitions...)
}

func (m *SuperFSM) Trigger(event IEvent) {
	for _, transition := range m.transitions {
		if reflect.TypeOf(transition.event) == reflect.TypeOf(event) && reflect.TypeOf(transition.from) == reflect.TypeOf(m.currentState()) && transition.guard(event) {
			m.SetState(transition.to, event)
			transition.action(event)
			break
		}
	}

	m.currentState().Trigger(event)
}

func (m *SuperFSM) currentState() IState {
	return m.states[m.currentStateIdx]
}
