package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type RootFSM struct {
	bot *Bot
	libs.SuperFSM[IBotState]
	UnimplementedBotState
}

func NewRootFSM(bot *Bot, states []IBotState, transitions []libs.Transition[IBotState]) *RootFSM {
	fsm := &RootFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM[IBotState](&NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

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
