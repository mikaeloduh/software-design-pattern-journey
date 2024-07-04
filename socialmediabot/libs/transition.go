package libs

// Transition
type Transition struct {
	from   IState
	to     IState
	event  IEvent
	guard  Guard
	action Action
}

func NewTransition(from IState, to IState, event IEvent, guard Guard, action Action) Transition {
	return Transition{from: from, to: to, event: event, guard: guard, action: action}
}

//func (t *Transition) Do(state IState, event IEvent) {
//	if state == t.from && event == t.event && t.guard() {
//		t.action()
//	}
//}
