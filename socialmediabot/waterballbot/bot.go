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

	normalStateFSM := NewNormalStateFSM(bot, &NullState{})
	normalStateFSM.AddState(defaultConversationState)
	normalStateFSM.AddState(interactingState)
	normalStateFSM.AddTransition(&NullState{}, defaultConversationState, EnterNormalStateEvent{}, EnterDefaultConversationGuard, NoAction)
	normalStateFSM.AddTransition(&NullState{}, interactingState, EnterNormalStateEvent{}, EnterInteractingGuard, NoAction)
	normalStateFSM.AddTransition(defaultConversationState, interactingState, service.NewLoginEvent{}, LoginEventGuard, NoAction)
	normalStateFSM.AddTransition(interactingState, defaultConversationState, service.NewLogoutEvent{}, LogoutEventGuard, NoAction)

	waitingState := NewWaitingState(bot)
	recordingState := NewRecordingState(bot)

	recordStateFSM := NewRecordStateFSM(bot, &NullState{})
	recordStateFSM.AddState(waitingState)
	recordStateFSM.AddState(recordingState)
	recordStateFSM.AddTransition(&NullState{}, waitingState, libs.EnterStateEvent{}, PositiveGuard, NoAction)
	recordStateFSM.AddTransition(waitingState, recordingState, service.GoBroadcastingEvent{}, PositiveGuard, NoAction)
	recordStateFSM.AddTransition(recordingState, waitingState, ExitRecordingStateEvent{}, PositiveGuard, NoAction)

	questioningState := NewQuestioningState(bot)
	thanksForJoiningState := NewThanksForJoiningState(bot)

	knowledgeKingStateFSM := NewKnowledgeKingStateFSM(bot, &NullState{})
	knowledgeKingStateFSM.AddState(questioningState)
	knowledgeKingStateFSM.AddState(thanksForJoiningState)
	knowledgeKingStateFSM.AddTransition(&NullState{}, questioningState, libs.EnterStateEvent{}, PositiveGuard, NoAction)
	knowledgeKingStateFSM.AddTransition(questioningState, thanksForJoiningState, ExitQuestioningStateEvent{}, PositiveGuard, NoAction)

	rootFSM := NewRootFSM(bot, &NullState{})
	rootFSM.AddState(normalStateFSM)
	rootFSM.AddState(recordStateFSM)
	rootFSM.AddState(knowledgeKingStateFSM)
	rootFSM.AddTransition(&NullState{}, normalStateFSM, libs.EnterStateEvent{}, PositiveGuard, NoAction)
	rootFSM.AddTransition(normalStateFSM, recordStateFSM, service.TagEvent{}, RecordCommandGuard, SaveCurrentRecorderAction)
	rootFSM.AddTransition(recordStateFSM, normalStateFSM, service.TagEvent{}, StopRecordCommandGuard, ClearCurrentRecorderAction)
	rootFSM.AddTransition(normalStateFSM, knowledgeKingStateFSM, service.TagEvent{}, KingCommandGuard, NoAction)
	rootFSM.AddTransition(knowledgeKingStateFSM, normalStateFSM, service.TagEvent{}, KingStopCommandGuard, NoAction)
	rootFSM.AddTransition(knowledgeKingStateFSM, normalStateFSM, ExitThanksForJoiningStateEvent{}, PositiveGuard, NoAction)

	bot.fsm = rootFSM

	rootFSM.Enter(nil)

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
