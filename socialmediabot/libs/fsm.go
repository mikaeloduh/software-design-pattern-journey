package libs

// IState
type IState interface {
}

type NormalState struct {
}

type Event string

type Guard func() bool

type Action func()

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

// FiniteStateMachine
type FiniteStateMachine struct {
	initState    IState
	currentState IState
	transitions  []Transition
}

func (m *FiniteStateMachine) GetState() IState {
	return m.currentState
}

func NewFiniteStateMachine(initState IState) *FiniteStateMachine {
	return &FiniteStateMachine{initState: initState, currentState: initState}
}
