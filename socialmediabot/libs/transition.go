package libs

// Transition
type Transition[T IState[T]] struct {
	from   T
	to     T
	event  IEvent
	guard  Guard
	action Action
}

func NewTransition[T IState[T]](from T, to T, event IEvent, guard Guard, action Action) Transition[T] {
	return Transition[T]{from: from, to: to, event: event, guard: guard, action: action}
}
