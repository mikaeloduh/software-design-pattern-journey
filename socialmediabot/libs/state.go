package libs

type SuperState[U any, T IFSM[T]] struct {
}

func (s *SuperState[U, T]) GetState() T {
	panic("to be implement")
}
