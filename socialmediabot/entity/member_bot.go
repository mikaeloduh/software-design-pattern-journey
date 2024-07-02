package entity

import "socialmediabot/libs"

// Bot
type Bot struct {
	id      string
	updater IUpdater
	fsm     *RootFSM
}

type BotGuard struct {
}

func (g BotGuard) Exec(event libs.IEvent) bool {
	return event.GetData().(NewLoginEvent).OnlineCount >= 10
}

func NewBot(waterball *Waterball) *Bot {
	bot := &Bot{id: "bot_001"}

	defaultConversationState := NewDefaultConversationState(waterball, bot)
	interactingState := NewInteractingState(waterball, bot)

	normalStateFSM := NewNormalStateFSM(waterball, bot, defaultConversationState)
	normalStateFSM.AddState(interactingState)
	normalStateFSM.AddTransition(libs.NewTransition(
		&DefaultConversationState{},
		&InteractingState{},
		NewLoginEvent{},
		BotGuard{},
		libs.Action(func() {}),
	))

	rootFSM := NewRootFSM(bot, normalStateFSM)
	bot.fsm = rootFSM

	return bot
}

func (b *Bot) Tag(event TagEvent) {
	//if event.TaggedTo == b {
	switch event.Message.Content {
	case "record":

	}
	//}
}

func (b *Bot) UpdateChatRoom(event NewMessageEvent) {
	b.fsm.OnNewMessage(event)
}

func (b *Bot) Update(event libs.IEvent) {
	switch value := event.(type) {
	case NewMessageEvent:
		if value.Sender == b {
			return
		}
		b.OnNewMessage(value)
	}

	b.fsm.Trigger(event)
}

func (b *Bot) NewMessageUpdate(event NewMessageEvent) {
	b.fsm.Trigger(event)
}

func (b *Bot) OnNewMessage(event NewMessageEvent) {
	b.fsm.OnNewMessage(event)
}

func (b *Bot) Id() string {
	return b.id
}
