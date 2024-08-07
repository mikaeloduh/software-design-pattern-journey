package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

// NormalStateFSM
type NormalStateFSM struct {
	bot *Bot
	libs.SuperFSM
	UnimplementedBotState
}

func NewNormalStateFSM(bot *Bot, states []libs.IState, transitions []libs.Transition) *NormalStateFSM {
	fsm := &NormalStateFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM(&NullState{}),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm
}

func (s *NormalStateFSM) Enter(_ libs.IEvent) {
	s.Trigger(EnterNormalStateEvent{OnlineCount: s.bot.waterball.OnlineCount()})
}

func (s *NormalStateFSM) Exit() {
	s.SetState(&NullState{}, nil)
}

func (s *NormalStateFSM) OnNewMessage(event service.NewMessageEvent) {
	s.GetState().(IBotState).OnNewMessage(event)
}

func (s *NormalStateFSM) OnNewPost(event service.NewPostEvent) {
	s.GetState().(IBotState).OnNewPost(event)
}

// EnterNormalStateEvent
type EnterNormalStateEvent struct {
	OnlineCount int
}

func (e EnterNormalStateEvent) GetData() libs.IEvent {
	return e
}
