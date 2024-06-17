package entity

import "socialmediabot/libs"

// NormalState
type NormalState struct {
	waterball *Waterball
	libs.SuperState[*Bot]
	UnimplementedBotState
}

func (s *NormalState) GetState() libs.IState {
	return s
}

func (s *NormalState) OnNewMessage(event NewMessageEvent) {
	s.waterball.ChatRoom.Send(s.Subject, Message{Content: "good to hear"})
}
