package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type WaitingState struct {
	bot *Bot
	libs.SuperState
	UnimplementedBotState
}

func NewWaitingState(bot *Bot) *WaitingState {
	return &WaitingState{
		bot:        bot,
		SuperState: libs.SuperState{},
	}
}

func (s *WaitingState) GetState() libs.IState {
	return s
}

func (s *WaitingState) OnNewMessage(_ service.NewMessageEvent) {
	// do nothing
}
