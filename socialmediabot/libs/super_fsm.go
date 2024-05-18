package libs

// SuperFSM
type SuperFSM interface {
}

// IState
type IState interface {
}

type Event string

type Guard func() bool

type Action func()
