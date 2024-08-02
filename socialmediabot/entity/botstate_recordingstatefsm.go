package entity

import "socialmediabot/libs"

type RecordStateFSM struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperFSM
	UnimplementedBotState
}

func NewRecordStateFSM(waterball *Waterball, bot *Bot, states []libs.IState, transitions []libs.Transition) *RecordStateFSM {
	fsm := &RecordStateFSM{
		bot:       bot,
		waterball: waterball,
		SuperFSM:  libs.NewSuperFSM(&NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (f *RecordStateFSM) Exit() {
	f.SetState(&NullState{}, nil)
}

func (f *RecordStateFSM) OnNewMessage(event NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}

func (f *RecordStateFSM) OnSpeak(event SpeakEvent) {
	f.GetState().(IBotState).OnSpeak(event)
}

func (f *RecordStateFSM) OnBroadcastStop(event BroadcastStopEvent) {
	f.GetState().(IBotState).OnBroadcastStop(event)
}
