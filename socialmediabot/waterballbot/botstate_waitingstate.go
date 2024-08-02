package waterballbot

import (
	"socialmediabot/entity"
	"socialmediabot/libs"
)

type WaitingState struct {
	bot       *Bot
	waterball *entity.Waterball
	libs.SuperState
	UnimplementedBotState
}

func NewWaitingState(waterball *entity.Waterball, bot *Bot) *WaitingState {
	return &WaitingState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *WaitingState) GetState() libs.IState {
	return s
}

func (s *WaitingState) OnNewMessage(_ entity.NewMessageEvent) {
	// do nothing
}
