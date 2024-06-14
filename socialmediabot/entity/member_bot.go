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
	UnimplementedBotState
	libs.SuperState[*Bot]
}

func (s *NormalState) GetState() libs.IFSM {
	return s
}

func (s *NormalState) OnNewMessage(event NewMessageEvent) {
	s.waterball.ChatRoom.Send(s.Subject, Message{Content: "good to hear"})
}

// Bot
type Bot struct {
	id      string
	updater IUpdater
	fsm     *libs.FiniteStateMachine[*Bot]
}

func NewBot(waterball *Waterball) *Bot {
	bot := Bot{id: "bot_001"}
	bot.fsm = libs.NewFiniteStateMachine[*Bot](&bot, &NormalState{
		waterball:  waterball,
		SuperState: libs.SuperState[*Bot]{Subject: &bot},
	})

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
		b.fsm.GetState().(IBotState).OnNewMessage(value)

	default:
		panic("unknown event")
	}

}

func (b *Bot) Id() string {
	return b.id
}
