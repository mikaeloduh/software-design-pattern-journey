package entity

import "socialmediabot/libs"

type RecordStateFSM struct {
	waterball *Waterball
	libs.SuperFSM[*Bot]
	UnimplementedBotState
}

func NewRecordStateFSM(waterball *Waterball, bot *Bot, initState libs.IState) *RecordStateFSM {
	return &RecordStateFSM{
		waterball: waterball,
		SuperFSM: libs.SuperFSM[*Bot]{
			Subject: bot,
			States:  []libs.IState{initState},
		},
	}
}

func (f *RecordStateFSM) OnNewMessage(event NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}
