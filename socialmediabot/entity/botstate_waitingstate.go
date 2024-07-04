package entity

import "socialmediabot/libs"

type WaitingState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
}

func NewWaitingState(waterball *Waterball, bot *Bot) *WaitingState {
	return &WaitingState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *WaitingState) GetState() libs.IState {
	return s
}

func (s *WaitingState) OnNewMessage(_ NewMessageEvent) {
	// do nothing
}
