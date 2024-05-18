package libs

import "socialmediabot/utils"

// Transition
type Transition[T IState] struct {
	from   T
	to     T
	event  Event
	guard  Guard
	action Action
}

func NewTransition[T IState](from T, to T, event Event, guard Guard, action Action) *Transition[T] {
	return &Transition[T]{from: from, to: to, event: event, guard: guard, action: action}
}

// FiniteStateMachine
type FiniteStateMachine[U any, T IState] struct {
	subject         U
	states          []T
	currentStateIdx int
	transitions     []Transition[T]
}

func NewFiniteStateMachine[U any, T IState](subject U, initState T) *FiniteStateMachine[U, T] {
	return &FiniteStateMachine[U, T]{
		subject:         subject,
		states:          []T{initState},
		currentStateIdx: 0,
	}
}

func (m *FiniteStateMachine[U, T]) GetSubject() U {
	return m.subject
}

func (m *FiniteStateMachine[U, T]) GetState() T {
	return m.states[m.currentStateIdx]
}

func (m *FiniteStateMachine[U, T]) SetState(state T) {
	for i, s := range m.states {
		if utils.ObjectsAreEqual(s, state) {
			m.currentStateIdx = i
			break
		}
	}
}

func (m *FiniteStateMachine[U, T]) AddState(state T) {
	m.states = append(m.states, state)
}

func (m *FiniteStateMachine[U, T]) AddTransition(transition *Transition[T]) {
	m.transitions = append(m.transitions, *transition)
}

func (m *FiniteStateMachine[U, T]) Trigger(event Event) {
	for _, transition := range m.transitions {
		if transition.event == event && utils.ObjectsAreEqual(transition.from, m.GetState()) && transition.guard() {
			m.SetState(transition.to)
			transition.action()
			break
		}
	}
}
