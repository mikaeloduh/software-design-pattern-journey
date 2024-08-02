package libs

// IState
type IState interface {
	Enter(event IEvent)
	Exit()
	GetState() IState
	Trigger(event IEvent)
}
