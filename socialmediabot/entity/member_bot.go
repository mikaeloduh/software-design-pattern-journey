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

	questioningState := NewQuestioningState(waterball, bot)
	thanksForJoiningState := NewThanksForJoiningState(waterball, bot)
	knowledgeKingStateFSM := NewKnowledgeKingStateFSM(waterball, bot,
		[]libs.IState{
			questioningState,
			thanksForJoiningState,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, questioningState, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(questioningState, thanksForJoiningState, libs.ExitStateEvent{}, PositiveGuard, NoAction),
		})

	rootFSM := NewRootFSM(bot,
		[]libs.IState{
			normalStateFSM,
			recordStateFSM,
			knowledgeKingStateFSM,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, normalStateFSM, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(normalStateFSM, recordStateFSM, TagEvent{}, RecordCommandGuard, NoAction),
			libs.NewTransition(normalStateFSM, knowledgeKingStateFSM, TagEvent{}, KingCommandGuard, NoAction),
			libs.NewTransition(recordStateFSM, normalStateFSM, TagEvent{}, StopRecordCommandGuard, NoAction),
		},
	)
	rootFSM.Enter()

	bot.fsm = rootFSM

	return bot
}

func KingCommandGuard(event libs.IEvent) bool {
	return event.GetData().(TagEvent).Message.Content == "king"
}

func PositiveGuard(event libs.IEvent) bool {
	return true
}

func (b *Bot) Tag(event TagEvent) {
	b.Update(event)
}

func (b *Bot) Update(event libs.IEvent) {
	switch value := event.(type) {
	case NewMessageEvent:
		if value.Sender == b {
			return
		}
		b.OnNewMessage(value)

	case NewPostEvent:
		b.OnNewPost(value)

	case TagEvent:
		b.fsm.Trigger(value)

	default:
		b.fsm.Trigger(value)
	}
}

func (b *Bot) OnNewMessage(event NewMessageEvent) {
	b.fsm.OnNewMessage(event)
}

func (b *Bot) OnNewPost(event NewPostEvent) {
	b.fsm.OnNewPost(event)
}

func (b *Bot) Id() string {
	return b.id
}

func (b *Bot) GetState() libs.IState {
	return b.fsm.GetState()
}
