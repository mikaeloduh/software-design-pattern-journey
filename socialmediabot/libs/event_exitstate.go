package libs

type ExitStateEvent struct {
}

func (e ExitStateEvent) GetData() IEvent {
	return e
}
