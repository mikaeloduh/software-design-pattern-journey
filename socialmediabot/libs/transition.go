package libs

// Transition
type Transition[T any] struct {
	from   T
	to     T
	event  Event
	guard  Guard
	action Action
}

func NewTransition[T IFSM[T]](from T, to T, event Event, guard Guard, action Action) *Transition[T] {
	return &Transition[T]{from: from, to: to, event: event, guard: guard, action: action}
}
