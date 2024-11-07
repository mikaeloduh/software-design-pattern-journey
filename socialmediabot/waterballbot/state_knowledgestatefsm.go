package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type KnowledgeKingStateFSM struct {
	bot *Bot
	libs.SuperFSM[IBotState]
	UnimplementedBotState
}

func NewKnowledgeKingStateFSM(bot *Bot, states []IBotState, transitions []libs.Transition[IBotState]) *KnowledgeKingStateFSM {
	fsm := &KnowledgeKingStateFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM[IBotState](&NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (f *KnowledgeKingStateFSM) Exit() {
	f.SetState(&NullState{}, nil)
}

func (f *KnowledgeKingStateFSM) OnSpeak(event service.SpeakEvent) {
	f.GetState().(IBotState).OnSpeak(event)
}

func (f *KnowledgeKingStateFSM) OnBroadcastStop(event service.BroadcastStopEvent) {
	f.GetState().(IBotState).OnBroadcastStop(event)
}
