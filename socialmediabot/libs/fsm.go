package libs

// IState
type IState interface {
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

func (m *FiniteStateMachine) AddTransition(transition *Transition) {
	m.transitions = append(m.transitions, *transition)
}

func (m *FiniteStateMachine) Trigger(event Event) {
	for _, transition := range m.transitions {
		if transition.from == m.currentState && transition.event == event && transition.guard() {
			m.currentState = transition.to
			transition.action()
			break
		}
	}
}

func NewFiniteStateMachine(initState IState) *FiniteStateMachine {
	return &FiniteStateMachine{initState: initState, currentState: initState}
}
