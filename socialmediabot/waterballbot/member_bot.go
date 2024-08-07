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
		[]libs.IState{
			defaultConversationState,
			interactingState,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, defaultConversationState, EnterNormalStateEvent{}, EnterDefaultConversationGuard, NoAction),
			libs.NewTransition(&NullState{}, interactingState, EnterNormalStateEvent{}, EnterInteractingGuard, NoAction),
			libs.NewTransition(defaultConversationState, interactingState, service.NewLoginEvent{}, LoginEventGuard, NoAction),
			libs.NewTransition(interactingState, defaultConversationState, service.NewLogoutEvent{}, LogoutEventGuard, NoAction),
		},
	)

	waitingState := NewWaitingState(bot)
	recordingState := NewRecordingState(bot)
	recordStateFSM := NewRecordStateFSM(bot,
		[]libs.IState{
			waitingState,
			recordingState,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, waitingState, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(waitingState, recordingState, service.GoBroadcastingEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(recordingState, waitingState, ExitRecordingStateEvent{}, PositiveGuard, NoAction),
		},
	)

	questioningState := NewQuestioningState(bot)
	thanksForJoiningState := NewThanksForJoiningState(bot)
	knowledgeKingStateFSM := NewKnowledgeKingStateFSM(bot,
		[]libs.IState{
			questioningState,
			thanksForJoiningState,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, questioningState, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(questioningState, thanksForJoiningState, ExitQuestioningStateEvent{}, PositiveGuard, NoAction),
		})

	rootFSM := NewRootFSM(bot,
		[]libs.IState{
			normalStateFSM,
			recordStateFSM,
			knowledgeKingStateFSM,
		},
		[]libs.Transition{
			libs.NewTransition(&NullState{}, normalStateFSM, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(normalStateFSM, recordStateFSM, service.TagEvent{}, RecordCommandGuard, SaveCurrentRecorderAction),
			libs.NewTransition(recordStateFSM, normalStateFSM, service.TagEvent{}, StopRecordCommandGuard, ClearCurrentRecorderAction),
			libs.NewTransition(normalStateFSM, knowledgeKingStateFSM, service.TagEvent{}, KingCommandGuard, NoAction),
			libs.NewTransition(knowledgeKingStateFSM, normalStateFSM, service.TagEvent{}, KingStopCommandGuard, NoAction),
			libs.NewTransition(knowledgeKingStateFSM, normalStateFSM, ExitThanksForJoiningStateEvent{}, PositiveGuard, NoAction),
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

func (b *Bot) GetState() libs.IState {
	return b.fsm.GetState()
}
