package entity

import "socialmediabot/libs"

type BotFSM struct {
	libs.SuperFSM[*Bot]
	UnimplementedBotState
}

func (f *BotFSM) OnNewMessage(event NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}
