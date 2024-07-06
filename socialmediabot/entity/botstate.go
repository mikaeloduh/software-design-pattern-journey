package entity

// IBotState
type IBotState interface {
	OnNewMessage(event NewMessageEvent)
	OnNewPost(event NewPostEvent)
}

type UnimplementedBotState struct{}

func (UnimplementedBotState) OnNewMessage(event NewMessageEvent) {
	panic("implement me")
}

func (UnimplementedBotState) OnNewPost(event NewPostEvent) {
	panic("implement me")
}
