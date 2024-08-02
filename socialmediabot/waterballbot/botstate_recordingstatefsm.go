package waterballbot

import (
	"socialmediabot/entity"
	"socialmediabot/libs"
)

type RecordStateFSM struct {
	bot       *Bot
	waterball *entity.Waterball
	libs.SuperFSM
	UnimplementedBotState
}

func NewRecordStateFSM(waterball *entity.Waterball, bot *Bot, states []libs.IState, transitions []libs.Transition) *RecordStateFSM {
	fsm := &RecordStateFSM{
		bot:       bot,
		waterball: waterball,
		SuperFSM:  libs.NewSuperFSM(&entity.NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (f *RecordStateFSM) Exit() {
	f.SetState(&entity.NullState{}, nil)
}

func (f *RecordStateFSM) OnNewMessage(event entity.NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}

func (f *RecordStateFSM) OnSpeak(event entity.SpeakEvent) {
	f.GetState().(IBotState).OnSpeak(event)
}

func (f *RecordStateFSM) OnBroadcastStop(event entity.BroadcastStopEvent) {
	f.GetState().(IBotState).OnBroadcastStop(event)
}
