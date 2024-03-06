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
type FiniteStateMachine[T IState] struct {
	initState    T
	currentState T
	transitions  []Transition[T]
}

func (m *FiniteStateMachine[T]) GetState() T {
	return m.currentState
}

func (m *FiniteStateMachine[T]) AddTransition(transition *Transition[T]) {
	m.transitions = append(m.transitions, *transition)
}

func (m *FiniteStateMachine[T]) Trigger(event Event) {
	for _, transition := range m.transitions {
		if utils.ObjectsAreEqual(transition.from, m.currentState) && transition.event == event && transition.guard() {
			m.currentState = transition.to
			transition.action()
			break
		}
	}
}

func NewFiniteStateMachine[T IState](initState T) *FiniteStateMachine[T] {
	return &FiniteStateMachine[T]{initState: initState, currentState: initState}
}
