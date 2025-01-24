package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

// IBotState
type IBotState interface {
	libs.IState[IBotState]
	IBotOperation
}

// IBotOperation
type IBotOperation interface {
	OnNewMessage(event service.NewMessageEvent)
	OnNewPost(event service.NewPostEvent)
	OnSpeak(event service.SpeakEvent)
	OnBroadcastStop(event service.BroadcastStopEvent)
}

type UnimplementedBotOperation struct{}

func (UnimplementedBotOperation) OnNewMessage(event service.NewMessageEvent) {
	panic("implement me")
}

func (UnimplementedBotOperation) OnNewPost(event service.NewPostEvent) {
	panic("implement me")
}

func (UnimplementedBotOperation) OnSpeak(event service.SpeakEvent) {
	panic("implement me")
}

func (UnimplementedBotOperation) OnBroadcastStop(event service.BroadcastStopEvent) {
	panic("implement me")
}
