package libs

type EnterStateEvent struct {
}

func (e EnterStateEvent) GetData() IEvent {
	return e
}
