package entity

import "socialmediabot/libs"

// NormalStateFSM
type NormalStateFSM struct {
	waterball *Waterball
	libs.SuperFSM[*Bot]
	UnimplementedBotState
}

func NewNormalStateFSM(waterball *Waterball, bot *Bot, initState libs.IState) *NormalStateFSM {
	return &NormalStateFSM{waterball: waterball, SuperFSM: libs.SuperFSM[*Bot]{
		Subject: bot,
		States:  []libs.IState{initState},
	}}
}

func (s *NormalStateFSM) OnNewMessage(event NewMessageEvent) {
	s.GetState().(IBotState).OnNewMessage(event)
}
