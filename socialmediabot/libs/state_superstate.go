package libs

type SuperState[U any] struct {
	Subject U
}

func (s *SuperState[U]) GetState() IState {
	panic("Unimplemented method")
}
