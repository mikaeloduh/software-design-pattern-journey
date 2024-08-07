package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type RootFSM struct {
	bot *Bot
	libs.SuperFSM
	UnimplementedBotState
}

func NewRootFSM(bot *Bot, states []libs.IState, transitions []libs.Transition) *RootFSM {
	fsm := &RootFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM(&NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (f *RootFSM) OnNewMessage(event service.NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}

func (f *RootFSM) OnNewPost(event service.NewPostEvent) {
	f.GetState().(IBotState).OnNewPost(event)
}

func (f *RootFSM) OnSpeak(event service.SpeakEvent) {
	f.GetState().(IBotState).OnSpeak(event)
}

func (f *RootFSM) OnBroadcastStop(event service.BroadcastStopEvent) {
	f.GetState().(IBotState).OnBroadcastStop(event)
}
