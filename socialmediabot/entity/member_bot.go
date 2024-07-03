package entity

import "socialmediabot/libs"

// Bot
type Bot struct {
	id  string
	fsm *RootFSM
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
		LoginEventGuard,
		NoAction,
	))

	waitingState := NewWaitingState(waterball, bot)
	recordingState := NewRecordingState(waterball, bot)
	recordStateFSM := NewRecordStateFSM(waterball, bot, waitingState)
	recordStateFSM.AddState(recordingState)

	rootFSM := NewRootFSM(bot, normalStateFSM)
	rootFSM.AddState(recordStateFSM)
	rootFSM.AddTransition(libs.NewTransition(
		&NormalStateFSM{},
		&RecordStateFSM{},
		TagEvent{},
		RecordCommandGuard,
		NoAction,
	))
	rootFSM.AddTransition(libs.NewTransition(
		&RecordStateFSM{},
		&NormalStateFSM{},
		TagEvent{},
		StopRecordCommandGuard,
		NoAction,
	))

	bot.fsm = rootFSM

	return bot
}

func (b *Bot) Tag(event TagEvent) {
	b.fsm.Trigger(event)
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

func (b *Bot) OnNewMessage(event NewMessageEvent) {
	b.fsm.OnNewMessage(event)
}

func (b *Bot) Id() string {
	return b.id
}
