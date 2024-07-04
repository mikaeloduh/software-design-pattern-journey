package entity

import "socialmediabot/libs"

// NormalStateFSM
type NormalStateFSM struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperFSM
	UnimplementedBotState
}

func NewNormalStateFSM(waterball *Waterball, bot *Bot, states []libs.IState, transitions []libs.Transition) *NormalStateFSM {
	fsm := &NormalStateFSM{
		bot:       bot,
		waterball: waterball,
		SuperFSM:  libs.NewSuperFSM(states[0]),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	fsm.Trigger(EnterNormalStateEvent{OnlineCount: waterball.OnlineCount()})

	return fsm
}

func (s *NormalStateFSM) Enter() {
	s.Trigger(EnterNormalStateEvent{OnlineCount: s.waterball.OnlineCount()})
}

func (s *NormalStateFSM) Exit() {
	s.SetState(&NullState{})
}

func (s *NormalStateFSM) OnNewMessage(event NewMessageEvent) {
	s.GetState().(IBotState).OnNewMessage(event)
}

// EnterNormalStateEvent
type EnterNormalStateEvent struct {
	OnlineCount int
}

func (e EnterNormalStateEvent) GetData() libs.IEvent {
	return e
}
