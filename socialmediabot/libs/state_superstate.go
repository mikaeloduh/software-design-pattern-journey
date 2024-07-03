package libs

type SuperState[U any] struct {
	Subject U
}

func (s *SuperState[U]) Enter() {
	//TODO implement me
}

func (s *SuperState[U]) Exit() {
	//TODO implement me
}

func (s *SuperState[U]) GetState() IState {
	panic("Unimplemented method")
}

func (s *SuperState[U]) Trigger(event IEvent) {
	// Do nothing
}
