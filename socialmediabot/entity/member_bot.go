package entity

import "socialmediabot/libs"

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
		SuperFSM: *libs.NewSuperFSM[*Bot](&bot, &initState),
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
