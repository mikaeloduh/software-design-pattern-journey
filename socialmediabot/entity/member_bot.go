package entity

import "socialmediabot/libs"

type IBotState interface {
	OnNewMessage(event NewMessageEvent)
}

type NormalState struct {
	fsm       *libs.FiniteStateMachine[*Bot, IBotState]
	waterball Waterball
}

func (s *NormalState) OnNewMessage(event NewMessageEvent) {
	s.waterball.ChatRoom.Send(s.fsm.GetSubject(), Message{Content: "good to hear"})
}

type Bot struct {
	id      string
	updater IUpdater
	fsm     *libs.FiniteStateMachine[*Bot, IBotState]
}

func NewBot(waterball Waterball) *Bot {
	bot := &Bot{id: "bot_001"}
	initState := &NormalState{waterball: waterball}
	fsm := libs.NewFiniteStateMachine[*Bot, IBotState](bot, initState)
	bot.fsm = fsm
	initState.fsm = fsm

	return bot
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
		state := b.fsm.GetState()
		state.OnNewMessage(value)

	default:
		panic("unknown event")
	}

}

func (b *Bot) Id() string {
	return b.id
}
