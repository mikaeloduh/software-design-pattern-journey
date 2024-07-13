package entity

import "socialmediabot/libs"

type KnowledgeKingStateFSM struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperFSM
	UnimplementedBotState
}

func NewKnowledgeKingStateFSM(waterball *Waterball, bot *Bot, states []libs.IState, transitions []libs.Transition) *KnowledgeKingStateFSM {
	fsm := &KnowledgeKingStateFSM{
		bot:       bot,
		waterball: waterball,
		SuperFSM:  libs.NewSuperFSM(&NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (f *KnowledgeKingStateFSM) OnSpeak(event SpeakEvent) {
	f.GetState().(IBotState).OnSpeak(event)
}

func (f *KnowledgeKingStateFSM) OnBroadcastStop(event BroadcastStopEvent) {
	f.GetState().(IBotState).OnBroadcastStop(event)
}
