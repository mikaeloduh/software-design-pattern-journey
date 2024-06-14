package libs

// IFSM
type IFSM interface {
	GetState() IFSM
}

type IState any

type Event string

type Guard func() bool

type Action func()
