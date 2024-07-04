package entity

import "socialmediabot/libs"

type InteractingState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
}

func NewInteractingState(waterball *Waterball, bot *Bot) *InteractingState {
	return &InteractingState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *InteractingState) GetState() libs.IState {
	return s
}

func (s *InteractingState) OnNewMessage(event NewMessageEvent) {
	s.waterball.ChatRoom.Send(s.bot, Message{Content: "I like your idea!"})
}
