package entity

import "socialmediabot/libs"

type RootFSM struct {
	bot *Bot
	libs.SuperFSM
	UnimplementedBotState
}

func NewRootFSM(bot *Bot, states []libs.IState, transitions []libs.Transition) *RootFSM {
	fsm := &RootFSM{
		bot:      bot,
		SuperFSM: libs.NewSuperFSM(states[0]),
	}
	fsm.AddState(states...)
	fsm.AddTransition(transitions...)

	return fsm

}

func (f *RootFSM) OnNewMessage(event NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}
