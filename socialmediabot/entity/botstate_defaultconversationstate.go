package entity

import "socialmediabot/libs"

// DefaultConversationState
type DefaultConversationState struct {
	waterball *Waterball
	libs.SuperState[*Bot]
	UnimplementedBotState
}

func NewDefaultConversationState(waterball *Waterball, bot *Bot) *DefaultConversationState {
	return &DefaultConversationState{waterball: waterball, SuperState: libs.SuperState[*Bot]{Subject: bot}}
}

func (s *DefaultConversationState) GetState() libs.IState {
	return s
}

func (s *DefaultConversationState) OnNewMessage(event NewMessageEvent) {
	s.waterball.ChatRoom.Send(s.Subject, Message{Content: "good to hear"})
}
