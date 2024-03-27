package libs

import "socialmediabot/utils"

// IState
type IState interface {
}

type Event string

type Guard func() bool

type Action func()

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
	subject      U
	initState    T
	currentState T
	transitions  []Transition[T]
}

func (m *FiniteStateMachine[U, T]) GetSubject() U {
	return m.subject
}

func (m *FiniteStateMachine[U, T]) GetState() T {
	return m.currentState
}

func (m *FiniteStateMachine[U, T]) AddTransition(transition *Transition[T]) {
	m.transitions = append(m.transitions, *transition)
}

func (m *FiniteStateMachine[U, T]) Trigger(event Event) {
	for _, transition := range m.transitions {
		if utils.ObjectsAreEqual(transition.from, m.currentState) && transition.event == event && transition.guard() {
			m.currentState = transition.to
			transition.action()
			break
		}
	}
}

func NewFiniteStateMachine[U any, T IState](subject U, initState T) *FiniteStateMachine[U, T] {
	return &FiniteStateMachine[U, T]{
		subject:      subject,
		initState:    initState,
		currentState: initState,
	}
}
