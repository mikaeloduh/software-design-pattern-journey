package entity

// Bot
type Bot struct {
	id      string
	updater IUpdater
	fsm     *RootFSM
}

func NewBot(waterball *Waterball) *Bot {
	bot := Bot{id: "bot_001"}
	bot.fsm = NewRootFSM(&bot, NewNormalStateFSM(waterball, &bot, NewDefaultConversationState(waterball, &bot)))

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
