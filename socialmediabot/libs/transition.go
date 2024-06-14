package libs

// Transition
type Transition struct {
	from   IFSM
	to     IFSM
	event  Event
	guard  Guard
	action Action
}

func NewTransition(from IFSM, to IFSM, event Event, guard Guard, action Action) *Transition {
	return &Transition{from: from, to: to, event: event, guard: guard, action: action}
}
