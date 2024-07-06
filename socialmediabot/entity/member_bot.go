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
	normalStateFSM := NewNormalStateFSM(waterball, bot,
		[]libs.IState{
			defaultConversationState,
			interactingState,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, defaultConversationState, EnterNormalStateEvent{}, EnterDefaultConversationGuard, NoAction),
			libs.NewTransition(&NullState{}, interactingState, EnterNormalStateEvent{}, EnterInteractingGuard, NoAction),
			libs.NewTransition(defaultConversationState, interactingState, NewLoginEvent{}, LoginEventGuard, NoAction),
		},
	)

	waitingState := NewWaitingState(waterball, bot)
	recordingState := NewRecordingState(waterball, bot)
	recordStateFSM := NewRecordStateFSM(waterball, bot,
		[]libs.IState{
			waitingState,
			recordingState,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, waitingState, libs.EnterStateEvent{}, PositiveGuard, NoAction),
		},
	)

	rootFSM := NewRootFSM(bot,
		[]libs.IState{
			normalStateFSM,
			recordStateFSM,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, normalStateFSM, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(normalStateFSM, recordStateFSM, TagEvent{}, RecordCommandGuard, NoAction),
			libs.NewTransition(recordStateFSM, normalStateFSM, TagEvent{}, StopRecordCommandGuard, NoAction),
		},
	)
	rootFSM.Enter()

	bot.fsm = rootFSM

	return bot
}

func PositiveGuard(event libs.IEvent) bool {
	return true
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
