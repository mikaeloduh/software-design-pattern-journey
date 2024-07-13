package entity

// IBotState
type IBotState interface {
	OnNewMessage(event NewMessageEvent)
	OnNewPost(event NewPostEvent)
	OnSpeak(event SpeakEvent)
	OnBroadcastStop(event BroadcastStopEvent)
}

type UnimplementedBotState struct{}

func (UnimplementedBotState) OnNewMessage(event NewMessageEvent) {
	panic("implement me")
}

func (UnimplementedBotState) OnNewPost(event NewPostEvent) {
	panic("implement me")
}

func (s UnimplementedBotState) OnSpeak(event SpeakEvent) {
	panic("implement me")
}

func (UnimplementedBotState) OnBroadcastStop(event BroadcastStopEvent) {
	panic("implement me")
}
