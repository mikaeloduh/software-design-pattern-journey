package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type WaitingState struct {
	bot *Bot
	libs.SuperState[IBotState]
	UnimplementedBotState
}

func NewWaitingState(bot *Bot) *WaitingState {
	return &WaitingState{
		bot:        bot,
		SuperState: libs.SuperState[IBotState]{},
	}
}

func (s *WaitingState) GetState() IBotState {
	return s
}

func (s *WaitingState) OnNewMessage(_ service.NewMessageEvent) {
	// do nothing
}
