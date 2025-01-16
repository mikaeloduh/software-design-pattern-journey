package libs

// IState
type IState[T any] interface {
	Enter(event IEvent)
	Exit()
	GetState() T
	Trigger(event IEvent)
}
