package entity

// IBotState
type IBotState interface {
	OnNewMessage(event NewMessageEvent)
}

type UnimplementedBotState struct{}

func (UnimplementedBotState) OnNewMessage(event NewMessageEvent) {
	panic("implement me")
}
