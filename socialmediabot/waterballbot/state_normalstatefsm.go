package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

// NormalStateFSM
type NormalStateFSM struct {
	bot *Bot
	libs.SuperFSM[IBotState]
	UnimplementedBotOperation
}

func NewNormalStateFSM(bot *Bot, initialState IBotState) *NormalStateFSM {
	fsm := &NormalStateFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM[IBotState](initialState),
	}

	return fsm
}

func (s *NormalStateFSM) Enter(_ libs.IEvent) {
	s.Trigger(EnterNormalStateEvent{OnlineCount: s.bot.waterball.OnlineCount()})
}

func (s *NormalStateFSM) Exit() {
	s.SetState(&NullState{}, nil)
}

func (s *NormalStateFSM) OnNewMessage(event service.NewMessageEvent) {
	s.GetState().OnNewMessage(event)
}

func (s *NormalStateFSM) OnNewPost(event service.NewPostEvent) {
	s.GetState().OnNewPost(event)
}

// EnterNormalStateEvent
type EnterNormalStateEvent struct {
	OnlineCount int
}

func (e EnterNormalStateEvent) GetData() libs.IEvent {
	return e
}
