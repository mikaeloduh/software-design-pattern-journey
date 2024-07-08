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

func (s *ThanksForJoiningState) Enter() {
	s.waterball.Broadcast.GoBroadcasting(s.bot)
	s.waterball.Broadcast.Transmit(NewSpeak(s.bot, "The winner is member_008"))
}

func (s *ThanksForJoiningState) GetState() libs.IState {
	return s
}
