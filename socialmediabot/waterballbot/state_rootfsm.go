package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type RootFSM struct {
	bot *Bot
	libs.SuperFSM[IBotState]
	UnimplementedBotOperation
}

func NewRootFSM(bot *Bot, initialState IBotState) *RootFSM {
	fsm := &RootFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM[IBotState](initialState),
	}

	return fsm
}

func (f *RootFSM) OnNewMessage(event service.NewMessageEvent) {
	f.GetState().OnNewMessage(event)
}

func (f *RootFSM) OnNewPost(event service.NewPostEvent) {
	f.GetState().OnNewPost(event)
}

func (f *RootFSM) OnSpeak(event service.SpeakEvent) {
	f.GetState().OnSpeak(event)
}

func (f *RootFSM) OnBroadcastStop(event service.BroadcastStopEvent) {
	f.GetState().OnBroadcastStop(event)
}
