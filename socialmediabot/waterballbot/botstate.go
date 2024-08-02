package waterballbot

import "socialmediabot/entity"

// IBotState
type IBotState interface {
	OnNewMessage(event entity.NewMessageEvent)
	OnNewPost(event entity.NewPostEvent)
	OnSpeak(event entity.SpeakEvent)
	OnBroadcastStop(event entity.BroadcastStopEvent)
}

type UnimplementedBotState struct{}

func (UnimplementedBotState) OnNewMessage(event entity.NewMessageEvent) {
	panic("implement me")
}

func (UnimplementedBotState) OnNewPost(event entity.NewPostEvent) {
	panic("implement me")
}

func (UnimplementedBotState) OnSpeak(event entity.SpeakEvent) {
	panic("implement me")
}

func (UnimplementedBotState) OnBroadcastStop(event entity.BroadcastStopEvent) {
	panic("implement me")
}
