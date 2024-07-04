package entity

import "socialmediabot/libs"

// Bot
type Bot struct {
	id  string
	fsm *RootFSM
}

func NewBot(waterball *Waterball) *Bot {
	bot := &Bot{id: "bot_001"}

	normalStateFSM := NewNormalStateFSM(waterball, bot,
		[]libs.IState{
			&NullState{},
			NewDefaultConversationState(waterball, bot),
			NewInteractingState(waterball, bot),
		},
		[]libs.Transition{
			libs.NewTransition(
				&NullState{},
				&DefaultConversationState{},
				EnterNormalStateEvent{},
				EnterDefaultConversationGuard,
				NoAction,
			), libs.NewTransition(
				&NullState{},
				&InteractingState{},
				EnterNormalStateEvent{},
				EnterInteractingGuard,
				NoAction,
			), libs.NewTransition(
				&DefaultConversationState{},
				&InteractingState{},
				NewLoginEvent{},
				LoginEventGuard,
				NoAction,
			),
		},
	)

	recordStateFSM := NewRecordStateFSM(waterball, bot,
		[]libs.IState{
			NewWaitingState(waterball, bot),
			NewRecordingState(waterball, bot),
		},
		nil,
	)

	rootFSM := NewRootFSM(bot,
		[]libs.IState{normalStateFSM, recordStateFSM},
		[]libs.Transition{
			libs.NewTransition(
				&NormalStateFSM{},
				&RecordStateFSM{},
				TagEvent{},
				RecordCommandGuard,
				NoAction,
			),
			libs.NewTransition(
				&RecordStateFSM{},
				&NormalStateFSM{},
				TagEvent{},
				StopRecordCommandGuard,
				NoAction,
			),
		},
	)

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
