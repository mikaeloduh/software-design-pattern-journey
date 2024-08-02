package waterballbot

import (
	"socialmediabot/entity"
	"socialmediabot/libs"
)

type RootFSM struct {
	bot *Bot
	libs.SuperFSM
	UnimplementedBotState
}

func NewRootFSM(bot *Bot, states []libs.IState, transitions []libs.Transition) *RootFSM {
	fsm := &RootFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM(&entity.NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (f *RootFSM) OnNewMessage(event entity.NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}

func (f *RootFSM) OnNewPost(event entity.NewPostEvent) {
	f.GetState().(IBotState).OnNewPost(event)
}

func (f *RootFSM) OnSpeak(event entity.SpeakEvent) {
	f.GetState().(IBotState).OnSpeak(event)
}

func (f *RootFSM) OnBroadcastStop(event entity.BroadcastStopEvent) {
	f.GetState().(IBotState).OnBroadcastStop(event)
}
