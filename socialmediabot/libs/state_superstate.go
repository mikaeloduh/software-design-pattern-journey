package libs

type SuperState struct {
}

func (s *SuperState) Enter(event IEvent) {
	//TODO implement me
}

func (s *SuperState) Exit() {
	//TODO implement me
}

func (s *SuperState) GetState() IState {
	panic("Unimplemented method")
}

func (s *SuperState) Trigger(event IEvent) {
	// Do nothing
}
