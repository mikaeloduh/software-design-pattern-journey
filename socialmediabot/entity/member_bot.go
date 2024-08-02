package entity

import (
	"socialmediabot/libs"
)

// Bot
type Bot struct {
	id       string
	role     Role
	fsm      *RootFSM
	recorder IMember
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
			libs.NewTransition(interactingState, defaultConversationState, NewLogoutEvent{}, LogoutEventGuard, NoAction),
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
			libs.NewTransition(waitingState, recordingState, GoBroadcastingEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(recordingState, waitingState, ExitRecordingStateEvent{}, PositiveGuard, NoAction),
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
			libs.NewTransition(questioningState, thanksForJoiningState, ExitQuestioningStateEvent{}, PositiveGuard, NoAction),
		})

	saveRecorderAction := func(arg any) {
		member := arg.(TagEvent).TaggedBy.(IMember)
		bot.SetRecorder(member)
	}

	rootFSM := NewRootFSM(bot,
		[]libs.IState{
			normalStateFSM,
			recordStateFSM,
			knowledgeKingStateFSM,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, normalStateFSM, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(normalStateFSM, recordStateFSM, TagEvent{}, RecordCommandGuard, saveRecorderAction),
			libs.NewTransition(recordStateFSM, normalStateFSM, StopRecordCommandEvent{}, StopRecordCommandGuard, NoAction),
			libs.NewTransition(normalStateFSM, knowledgeKingStateFSM, TagEvent{}, KingCommandGuard, NoAction),
			libs.NewTransition(knowledgeKingStateFSM, normalStateFSM, ExitThanksForJoiningStateEvent{}, PositiveGuard, NoAction),
		},
	)
	rootFSM.Enter(nil)

	bot.fsm = rootFSM

	return bot
}

func (b *Bot) Tag(event TagEvent) {
	if event.Message.Content == "stop-recording" {
		b.Update(StopRecordCommandEvent{
			TaggedBy: event.TaggedBy,
			TaggedTo: event.TaggedTo,
			Message:  event.Message,
			Recorder: b.Recorder(),
		})
	} else {
		b.Update(event)
	}
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

	case SpeakEvent:
		b.OnSpeak(value)

	case BroadcastStopEvent:
		b.OnBroadcastStop(value)

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

func (b *Bot) OnSpeak(event SpeakEvent) {
	b.fsm.OnSpeak(event)
}

func (b *Bot) OnBroadcastStop(event BroadcastStopEvent) {
	b.fsm.OnBroadcastStop(event)
}

func (b *Bot) Id() string {
	return b.id
}

func (b *Bot) Role() Role {
	return b.role
}

func (b *Bot) Recorder() IMember {
	return b.recorder
}

func (b *Bot) SetRecorder(recorder IMember) {
	b.recorder = recorder
}

func (b *Bot) GetState() libs.IState {
	return b.fsm.GetState()
}
