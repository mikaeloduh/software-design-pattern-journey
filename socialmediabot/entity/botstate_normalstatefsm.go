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
		SuperFSM:  libs.NewSuperFSM(&NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (s *NormalStateFSM) Enter(_ libs.IEvent) {
	s.Trigger(EnterNormalStateEvent{OnlineCount: s.waterball.OnlineCount()})
}

func (s *NormalStateFSM) Exit() {
	s.SetState(&NullState{}, nil)
}

func (s *NormalStateFSM) OnNewMessage(event NewMessageEvent) {
	s.GetState().(IBotState).OnNewMessage(event)
}

func (s *NormalStateFSM) OnNewPost(event NewPostEvent) {
	s.GetState().(IBotState).OnNewPost(event)
}

// EnterNormalStateEvent
type EnterNormalStateEvent struct {
	OnlineCount int
}

func (e EnterNormalStateEvent) GetData() libs.IEvent {
	return e
}
