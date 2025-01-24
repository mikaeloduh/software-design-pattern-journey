package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type RecordStateFSM struct {
	bot *Bot
	libs.SuperFSM[IBotState]
	UnimplementedBotOperation
}

func NewRecordStateFSM(bot *Bot, initialState IBotState) *RecordStateFSM {
	fsm := &RecordStateFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM[IBotState](initialState),
	}

	return fsm
}

func (f *RecordStateFSM) Exit() {
	f.SetState(&NullState{}, nil)
}

func (f *RecordStateFSM) OnNewMessage(event service.NewMessageEvent) {
	f.GetState().OnNewMessage(event)
}

func (f *RecordStateFSM) OnSpeak(event service.SpeakEvent) {
	f.GetState().OnSpeak(event)
}

func (f *RecordStateFSM) OnBroadcastStop(event service.BroadcastStopEvent) {
	f.GetState().OnBroadcastStop(event)
}
