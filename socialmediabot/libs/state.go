package libs

type SuperState[U any] struct {
	Subject U
}

func (s *SuperState[U]) GetState() IFSM {
	panic("Unimplemented method")
}
