package waterballbot

import "socialmediabot/service"

// IBotState
type IBotState interface {
	OnNewMessage(event service.NewMessageEvent)
	OnNewPost(event service.NewPostEvent)
	OnSpeak(event service.SpeakEvent)
	OnBroadcastStop(event service.BroadcastStopEvent)
}

type UnimplementedBotState struct{}

func (UnimplementedBotState) OnNewMessage(event service.NewMessageEvent) {
	panic("implement me")
}

func (UnimplementedBotState) OnNewPost(event service.NewPostEvent) {
	panic("implement me")
}

func (UnimplementedBotState) OnSpeak(event service.SpeakEvent) {
	panic("implement me")
}

func (UnimplementedBotState) OnBroadcastStop(event service.BroadcastStopEvent) {
	panic("implement me")
}
