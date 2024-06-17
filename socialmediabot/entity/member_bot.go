package entity

import "socialmediabot/libs"

// IBotState
type IBotState interface {
	OnNewMessage(event NewMessageEvent)
}

type UnimplementedBotState struct{}

func (UnimplementedBotState) OnNewMessage(event NewMessageEvent) {
	panic("implement me")
}

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

type BotFSM struct {
	*libs.SuperFSM[*Bot]
	UnimplementedBotState
}

func (f *BotFSM) OnNewMessage(event NewMessageEvent) {
	f.GetState().(IBotState).OnNewMessage(event)
}

// Bot
type Bot struct {
	id      string
	updater IUpdater
	fsm     BotFSM
}

func NewBot(waterball *Waterball) *Bot {
	bot := Bot{id: "bot_001"}
	initState := NormalState{
		waterball:  waterball,
		SuperState: libs.SuperState[*Bot]{Subject: &bot},
	}
	bot.fsm = BotFSM{
		SuperFSM: libs.NewSuperFSM[*Bot](&bot, &initState),
	}

	return &bot
}

func (b *Bot) Tag(event TagEvent) {
	//TODO implement me
	panic("implement me")
}

func (b *Bot) Update(event NewMessageEvent) {
	b.updateHandler(event)
}

func (b *Bot) updateHandler(event IEvent) {
	switch value := event.(type) {
	case NewMessageEvent:
		if value.Sender == b {
			return
		}
		b.fsm.OnNewMessage(value)

	default:
		panic("unknown event")
	}

}

func (b *Bot) Id() string {
	return b.id
}
