package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type WaitingState struct {
	bot       *Bot
	waterball *service.Waterball
	libs.SuperState
	UnimplementedBotState
}

func NewWaitingState(waterball *service.Waterball, bot *Bot) *WaitingState {
	return &WaitingState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *WaitingState) GetState() libs.IState {
	return s
}

func (s *WaitingState) OnNewMessage(_ service.NewMessageEvent) {
	// do nothing
}
