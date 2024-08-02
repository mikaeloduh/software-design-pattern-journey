package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type KnowledgeKingStateFSM struct {
	bot       *Bot
	waterball *service.Waterball
	libs.SuperFSM
	UnimplementedBotState
}

func NewKnowledgeKingStateFSM(waterball *service.Waterball, bot *Bot, states []libs.IState, transitions []libs.Transition) *KnowledgeKingStateFSM {
	fsm := &KnowledgeKingStateFSM{
		bot:       bot,
		waterball: waterball,
		SuperFSM:  libs.NewSuperFSM(&NullState{}),
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
