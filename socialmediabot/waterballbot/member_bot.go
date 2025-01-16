package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

// Bot
type Bot struct {
	id        string
	role      service.Role
	fsm       *RootFSM
	waterball *service.Waterball
	recorder  service.IMember
}

func NewBot(waterball *service.Waterball) *Bot {
	bot := &Bot{
		id:        "bot_001",
		waterball: waterball,
	}

	defaultConversationState := NewDefaultConversationState(bot)
	interactingState := NewInteractingState(bot)
	normalStateFSM := NewNormalStateFSM(bot,
		[]IBotState{
			defaultConversationState,
			interactingState,
		},
		[]libs.Transition[IBotState]{
			libs.NewTransition[IBotState](&NullState{}, defaultConversationState, EnterNormalStateEvent{}, EnterDefaultConversationGuard, NoAction),
			libs.NewTransition[IBotState](&NullState{}, interactingState, EnterNormalStateEvent{}, EnterInteractingGuard, NoAction),
			libs.NewTransition[IBotState](defaultConversationState, interactingState, service.NewLoginEvent{}, LoginEventGuard, NoAction),
			libs.NewTransition[IBotState](interactingState, defaultConversationState, service.NewLogoutEvent{}, LogoutEventGuard, NoAction),
		},
	)

	waitingState := NewWaitingState(bot)
	recordingState := NewRecordingState(bot)
	recordStateFSM := NewRecordStateFSM(bot,
		[]IBotState{
			waitingState,
			recordingState,
		},
		[]libs.Transition[IBotState]{
			libs.NewTransition[IBotState](&NullState{}, waitingState, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition[IBotState](waitingState, recordingState, service.GoBroadcastingEvent{}, PositiveGuard, NoAction),
			libs.NewTransition[IBotState](recordingState, waitingState, ExitRecordingStateEvent{}, PositiveGuard, NoAction),
		},
	)

	questioningState := NewQuestioningState(bot)
	thanksForJoiningState := NewThanksForJoiningState(bot)
	knowledgeKingStateFSM := NewKnowledgeKingStateFSM(bot,
		[]IBotState{
			questioningState,
			thanksForJoiningState,
		},
		[]libs.Transition[IBotState]{
			libs.NewTransition[IBotState](&NullState{}, questioningState, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition[IBotState](questioningState, thanksForJoiningState, ExitQuestioningStateEvent{}, PositiveGuard, NoAction),
		})

	rootFSM := NewRootFSM(bot,
		[]IBotState{
			normalStateFSM,
			recordStateFSM,
			knowledgeKingStateFSM,
		},
		[]libs.Transition[IBotState]{
			libs.NewTransition[IBotState](&NullState{}, normalStateFSM, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition[IBotState](normalStateFSM, recordStateFSM, service.TagEvent{}, RecordCommandGuard, SaveCurrentRecorderAction),
			libs.NewTransition[IBotState](recordStateFSM, normalStateFSM, service.TagEvent{}, StopRecordCommandGuard, ClearCurrentRecorderAction),
			libs.NewTransition[IBotState](normalStateFSM, knowledgeKingStateFSM, service.TagEvent{}, KingCommandGuard, NoAction),
			libs.NewTransition[IBotState](knowledgeKingStateFSM, normalStateFSM, service.TagEvent{}, KingStopCommandGuard, NoAction),
			libs.NewTransition[IBotState](knowledgeKingStateFSM, normalStateFSM, ExitThanksForJoiningStateEvent{}, PositiveGuard, NoAction),
		},
	)
	rootFSM.Enter(nil)

	bot.fsm = rootFSM

	return bot
}

func (b *Bot) Tag(event service.TagEvent) {
	b.Update(event)
}

func (b *Bot) Update(event libs.IEvent) {
	switch value := event.(type) {
	case service.NewMessageEvent:
		if value.Sender == b {
			return
		}
		b.OnNewMessage(value)

	case service.NewPostEvent:
		b.OnNewPost(value)

	case service.SpeakEvent:
		b.OnSpeak(value)

	case service.BroadcastStopEvent:
		b.OnBroadcastStop(value)

	case service.TagEvent:
		b.fsm.Trigger(value)

	default:
		b.fsm.Trigger(value)
	}
}

func (b *Bot) OnNewMessage(event service.NewMessageEvent) {
	b.fsm.OnNewMessage(event)
}

func (b *Bot) OnNewPost(event service.NewPostEvent) {
	b.fsm.OnNewPost(event)
}

func (b *Bot) OnSpeak(event service.SpeakEvent) {
	b.fsm.OnSpeak(event)
}

func (b *Bot) OnBroadcastStop(event service.BroadcastStopEvent) {
	b.fsm.OnBroadcastStop(event)
}

func (b *Bot) Id() string {
	return b.id
}

func (b *Bot) Role() service.Role {
	return b.role
}

func (b *Bot) Recorder() service.IMember {
	return b.recorder
}

func (b *Bot) SetRecorder(recorder service.IMember) {
	b.recorder = recorder
}

func (b *Bot) GetState() IBotState {
	return b.fsm.GetState()
}
