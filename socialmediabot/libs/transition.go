package libs

// Transition
type Transition struct {
	from   IState
	to     IState
	event  Event
	guard  Guard
	action Action
}

func NewTransition(from IState, to IState, event Event, guard Guard, action Action) *Transition {
	return &Transition{from: from, to: to, event: event, guard: guard, action: action}
}
