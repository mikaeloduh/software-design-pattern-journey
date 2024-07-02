package entity

import "socialmediabot/libs"

type InteractingState struct {
	waterball *Waterball
	libs.SuperState[*Bot]
	UnimplementedBotState
}

func NewInteractingState(waterball *Waterball, bot *Bot) *InteractingState {
	return &InteractingState{waterball: waterball, SuperState: libs.SuperState[*Bot]{Subject: bot}}
}

func (s *InteractingState) GetState() libs.IState {
	return s
}

func (s *InteractingState) OnNewMessage(event NewMessageEvent) {
	s.waterball.ChatRoom.Send(s.Subject, Message{Content: "I like your idea!"})
}
