package entity

import "socialmediabot/libs"

type WaitingState struct {
	waterball *Waterball
	libs.SuperState[*Bot]
	UnimplementedBotState
}

func NewWaitingState(waterball *Waterball, bot *Bot) *WaitingState {
	return &WaitingState{
		waterball: waterball,
		SuperState: libs.SuperState[*Bot]{
			Subject: bot,
		},
	}
}

func (s *WaitingState) GetState() libs.IState {
	return s
}

func (s *WaitingState) OnNewMessage(_ NewMessageEvent) {
	// do nothing
}
