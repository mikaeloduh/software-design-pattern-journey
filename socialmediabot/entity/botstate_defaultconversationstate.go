package entity

import "socialmediabot/libs"

// DefaultConversationState
type DefaultConversationState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
}

func NewDefaultConversationState(waterball *Waterball, bot *Bot) *DefaultConversationState {
	return &DefaultConversationState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *DefaultConversationState) GetState() libs.IState {
	return s
}

func (s *DefaultConversationState) OnNewMessage(event NewMessageEvent) {
	s.waterball.ChatRoom.Send(s.bot, Message{Content: "good to hear"})
}
