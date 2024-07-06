package entity

import "socialmediabot/libs"

type ThanksForJoiningState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
}

func NewThanksForJoiningState(waterball *Waterball, bot *Bot) *ThanksForJoiningState {
	return &ThanksForJoiningState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *ThanksForJoiningState) GetState() libs.IState {
	return s
}
