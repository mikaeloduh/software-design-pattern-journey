package libs

// IFSM
type IFSM[T any] interface {
	GetState() T
}

type IState any

type Event string

type Guard func() bool

type Action func()
