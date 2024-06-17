package libs

import "socialmediabot/utils"

// IState

// SuperFSM is a finite state machine, U is the subject type, T is the state type
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
	}
}

func (m *SuperFSM[U]) GetSubject() U {
	return m.Subject
}

func (m *SuperFSM[U]) GetState() IState {
	return m.States[m.currentStateIdx].GetState()
}

func (m *SuperFSM[U]) SetState(state IState) {
	for i, s := range m.States {
		if utils.ObjectsAreEqual(s, state) {
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

func (m *SuperFSM[U]) Trigger(event Event) {
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
