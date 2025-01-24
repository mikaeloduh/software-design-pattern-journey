package libs

type SuperState[T any] struct {
}

func (s *SuperState[T]) Enter(event IEvent) {
	//TODO implement me
}

func (s *SuperState[T]) Exit() {
	//TODO implement me
}

func (s *SuperState[T]) GetState() T {
	panic("Unimplemented method")
}

func (s *SuperState[T]) Trigger(event IEvent) {
	// Do nothing
}
