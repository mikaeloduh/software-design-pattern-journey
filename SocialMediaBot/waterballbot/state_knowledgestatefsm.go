package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type KnowledgeKingStateFSM struct {
	bot *Bot
	libs.SuperFSM[IBotState]
	UnimplementedBotOperation
}

func NewKnowledgeKingStateFSM(bot *Bot, initialState IBotState) *KnowledgeKingStateFSM {
	fsm := &KnowledgeKingStateFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM[IBotState](initialState),
	}

	return fsm
}

func (f *KnowledgeKingStateFSM) Exit() {
	f.SetState(&NullState{}, nil)
}

func (f *KnowledgeKingStateFSM) OnSpeak(event service.SpeakEvent) {
	f.GetState().OnSpeak(event)
}

func (f *KnowledgeKingStateFSM) OnBroadcastStop(event service.BroadcastStopEvent) {
	f.GetState().OnBroadcastStop(event)
}
