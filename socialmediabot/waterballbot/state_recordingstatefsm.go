package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type RecordStateFSM struct {
	bot *Bot
	libs.SuperFSM
	UnimplementedBotState
}

func NewRecordStateFSM(bot *Bot, states []libs.IState, transitions []libs.Transition) *RecordStateFSM {
	fsm := &RecordStateFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM(&NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (f *RecordStateFSM) Exit() {
	f.SetState(&NullState{}, nil)
}

func (f *RecordStateFSM) OnNewMessage(event service.NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}

func (f *RecordStateFSM) OnSpeak(event service.SpeakEvent) {
	f.GetState().(IBotState).OnSpeak(event)
}

func (f *RecordStateFSM) OnBroadcastStop(event service.BroadcastStopEvent) {
	f.GetState().(IBotState).OnBroadcastStop(event)
}
