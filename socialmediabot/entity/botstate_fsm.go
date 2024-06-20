package entity

import "socialmediabot/libs"

type BotFSM struct {
	libs.SuperFSM[*Bot]
	UnimplementedBotState
}

func NewBotFSM(bot *Bot, initState libs.IState) *BotFSM {
	return &BotFSM{SuperFSM: *libs.NewSuperFSM[*Bot](bot, initState)}
}

func (f *BotFSM) OnNewMessage(event NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}
