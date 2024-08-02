package waterballbot

import (
	"socialmediabot/entity"
	"socialmediabot/libs"
)

// Bot
type Bot struct {
	id       string
	role     entity.Role
	fsm      *RootFSM
	recorder entity.IMember
}

func NewBot(waterball *entity.Waterball) *Bot {
	bot := &Bot{id: "bot_001"}

	defaultConversationState := NewDefaultConversationState(waterball, bot)
	interactingState := NewInteractingState(waterball, bot)
	normalStateFSM := NewNormalStateFSM(waterball, bot,
		[]libs.IState{
			defaultConversationState,
			interactingState,
		},
		[]libs.Transition{
			libs.NewTransition(&entity.NullState{}, defaultConversationState, EnterNormalStateEvent{}, EnterDefaultConversationGuard, NoAction),
			libs.NewTransition(&entity.NullState{}, interactingState, EnterNormalStateEvent{}, EnterInteractingGuard, NoAction),
			libs.NewTransition(defaultConversationState, interactingState, entity.NewLoginEvent{}, LoginEventGuard, NoAction),
			libs.NewTransition(interactingState, defaultConversationState, entity.NewLogoutEvent{}, LogoutEventGuard, NoAction),
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
			libs.NewTransition(&entity.NullState{}, waitingState, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(waitingState, recordingState, entity.GoBroadcastingEvent{}, PositiveGuard, NoAction),
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
			libs.NewTransition(&entity.NullState{}, questioningState, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(questioningState, thanksForJoiningState, ExitQuestioningStateEvent{}, PositiveGuard, NoAction),
		})

	saveRecorderAction := func(arg any) {
		member := arg.(entity.TagEvent).TaggedBy.(entity.IMember)
		bot.SetRecorder(member)
	}

	rootFSM := NewRootFSM(bot,
		[]libs.IState{
			normalStateFSM,
			recordStateFSM,
			knowledgeKingStateFSM,
		},
		[]libs.Transition{
			libs.NewTransition(&entity.NullState{}, normalStateFSM, libs.EnterStateEvent{}, PositiveGuard, NoAction),
			libs.NewTransition(normalStateFSM, recordStateFSM, entity.TagEvent{}, RecordCommandGuard, saveRecorderAction),
			libs.NewTransition(recordStateFSM, normalStateFSM, StopRecordCommandEvent{}, StopRecordCommandGuard, NoAction),
			libs.NewTransition(normalStateFSM, knowledgeKingStateFSM, entity.TagEvent{}, KingCommandGuard, NoAction),
			libs.NewTransition(knowledgeKingStateFSM, normalStateFSM, entity.TagEvent{}, KingStopCommandGuard, NoAction),
			libs.NewTransition(knowledgeKingStateFSM, normalStateFSM, ExitThanksForJoiningStateEvent{}, PositiveGuard, NoAction),
		},
	)
	rootFSM.Enter(nil)

	bot.fsm = rootFSM

	return bot
}

func (b *Bot) Tag(event entity.TagEvent) {
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
	case entity.NewMessageEvent:
		if value.Sender == b {
			return
		}
		b.OnNewMessage(value)

	case entity.NewPostEvent:
		b.OnNewPost(value)

	case entity.SpeakEvent:
		b.OnSpeak(value)

	case entity.BroadcastStopEvent:
		b.OnBroadcastStop(value)

	case entity.TagEvent:
		b.fsm.Trigger(value)

	default:
		b.fsm.Trigger(value)
	}
}

func (b *Bot) OnNewMessage(event entity.NewMessageEvent) {
	b.fsm.OnNewMessage(event)
}

func (b *Bot) OnNewPost(event entity.NewPostEvent) {
	b.fsm.OnNewPost(event)
}

func (b *Bot) OnSpeak(event entity.SpeakEvent) {
	b.fsm.OnSpeak(event)
}

func (b *Bot) OnBroadcastStop(event entity.BroadcastStopEvent) {
	b.fsm.OnBroadcastStop(event)
}

func (b *Bot) Id() string {
	return b.id
}

func (b *Bot) Role() entity.Role {
	return b.role
}

func (b *Bot) Recorder() entity.IMember {
	return b.recorder
}

func (b *Bot) SetRecorder(recorder entity.IMember) {
	b.recorder = recorder
}

func (b *Bot) GetState() libs.IState {
	return b.fsm.GetState()
}
