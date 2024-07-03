package entity

import "socialmediabot/libs"

// Bot
type Bot struct {
	id string
	//updater IUpdater
	fsm *RootFSM
}

type LoginGuard struct {
}

func (g LoginGuard) Exec(event libs.IEvent) bool {
	return event.GetData().(NewLoginEvent).OnlineCount >= 10
}

func LoginEventGuard(event libs.IEvent) bool {
	return event.GetData().(NewLoginEvent).OnlineCount >= 10
}

type RecordGuard struct {
}

func (g RecordGuard) Exec(event libs.IEvent) bool {
	return event.GetData().(NewMessageEvent).Message.Content == "record"
}

func RecordEventGuard(event libs.IEvent) bool {
	return event.GetData().(TagEvent).Message.Content == "record"
}

func NoAction() {}

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
		RecordEventGuard,
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
