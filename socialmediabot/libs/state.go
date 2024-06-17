package libs

// IState
type IState interface {
	GetState() IState
}

type Event string

type Guard func() bool

type Action func()
