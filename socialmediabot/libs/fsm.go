package libs

import "socialmediabot/utils"

// IFSM

// FiniteStateMachine is a finite state machine, U is the subject type, T is the state type
type FiniteStateMachine[U any] struct {
	subject         U
	states          []IFSM
	currentStateIdx int
	transitions     []Transition
}

func NewFiniteStateMachine[U any](subject U, initState IFSM) *FiniteStateMachine[U] {
	return &FiniteStateMachine[U]{
		subject:         subject,
		states:          []IFSM{initState},
		currentStateIdx: 0,
	}
}

func (m *FiniteStateMachine[U]) GetSubject() U {
	return m.subject
}

func (m *FiniteStateMachine[U]) GetState() IFSM {
	return m.states[m.currentStateIdx].GetState()
}

func (m *FiniteStateMachine[U]) SetState(state IFSM) {
	for i, s := range m.states {
		if utils.ObjectsAreEqual(s, state) {
			m.currentStateIdx = i
			break
		}
	}
}

func (m *FiniteStateMachine[U]) AddState(state IFSM) {
	m.states = append(m.states, state)
}

func (m *FiniteStateMachine[U]) AddTransition(transition *Transition) {
	m.transitions = append(m.transitions, *transition)
}

func (m *FiniteStateMachine[U]) Trigger(event Event) {
	for _, transition := range m.transitions {
		if transition.event == event && utils.ObjectsAreEqual(transition.from, m.GetState()) && transition.guard() {
			m.SetState(transition.to)
			transition.action()
			break
		}
	}
}

// Default FSM
type DefaultImplementationFSM struct {
}

func (f *DefaultImplementationFSM) GetState() {

}
