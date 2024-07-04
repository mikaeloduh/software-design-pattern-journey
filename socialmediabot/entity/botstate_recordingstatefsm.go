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
		SuperFSM:  libs.NewSuperFSM(states[0]),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (f *RecordStateFSM) OnNewMessage(event NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}
