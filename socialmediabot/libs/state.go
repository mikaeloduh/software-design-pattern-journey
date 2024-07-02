package libs

// IState
type IState interface {
	Enter()
	Exit()
	GetState() IState
	Trigger(event IEvent)
}
