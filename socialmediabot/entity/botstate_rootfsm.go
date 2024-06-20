package entity

import "socialmediabot/libs"

type RootFSM struct {
	libs.SuperFSM[*Bot]
	UnimplementedBotState
}

func NewRootFSM(bot *Bot, initState libs.IState) *RootFSM {
	return &RootFSM{SuperFSM: *libs.NewSuperFSM[*Bot](bot, initState)}
}

func (f *RootFSM) OnNewMessage(event NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}
