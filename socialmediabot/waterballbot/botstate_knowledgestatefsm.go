package waterballbot

import (
	"socialmediabot/entity"
	"socialmediabot/libs"
)

type KnowledgeKingStateFSM struct {
	bot       *Bot
	waterball *entity.Waterball
	libs.SuperFSM
	UnimplementedBotState
}

func NewKnowledgeKingStateFSM(waterball *entity.Waterball, bot *Bot, states []libs.IState, transitions []libs.Transition) *KnowledgeKingStateFSM {
	fsm := &KnowledgeKingStateFSM{
		bot:       bot,
		waterball: waterball,
		SuperFSM:  libs.NewSuperFSM(&entity.NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (f *KnowledgeKingStateFSM) Exit() {
	f.SetState(&entity.NullState{}, nil)
}

func (f *KnowledgeKingStateFSM) OnSpeak(event entity.SpeakEvent) {
	f.GetState().(IBotState).OnSpeak(event)
}

func (f *KnowledgeKingStateFSM) OnBroadcastStop(event entity.BroadcastStopEvent) {
	f.GetState().(IBotState).OnBroadcastStop(event)
}
