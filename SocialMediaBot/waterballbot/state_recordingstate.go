package waterballbot

import (
	"fmt"
	"socialmediabot/libs"
	"socialmediabot/service"
)

// RecordingState
type RecordingState struct {
	bot *Bot
	libs.SuperState[IBotState]
	UnimplementedBotOperation
	record *Record
}

func NewRecordingState(bot *Bot) *RecordingState {
	return &RecordingState{
		bot:        bot,
		SuperState: libs.SuperState[IBotState]{},
	}
}

func (s *RecordingState) GetState() IBotState {
	return s
}

func (s *RecordingState) Enter(_ libs.IEvent) {
	s.record = &Record{}
}

func (s *RecordingState) Exit() {
	s.replayRecord()
}

func (s *RecordingState) OnNewMessage(_ service.NewMessageEvent) {
}

func (s *RecordingState) OnSpeak(event service.SpeakEvent) {
	s.record.AddContent(event.Content)
}

func (s *RecordingState) OnBroadcastStop(_ service.BroadcastStopEvent) {
	s.bot.Update(ExitRecordingStateEvent{})
}

func (s *RecordingState) replayRecord() {
	s.bot.waterball.ChatRoom.Send(service.NewMessage(s.bot, fmt.Sprintf("[Record Replay] %s", s.record.GetContent())))
}

// Record
type Record struct {
	content string
}

func (r *Record) AddContent(content string) {
	if len(r.content) == 0 {
		r.content = content
	} else {
		r.content = r.content + "\n" + content
	}
}

func (r *Record) GetContent() string {
	return r.content
}

// ExitRecordingStateEvent
type ExitRecordingStateEvent struct {
}

func (e ExitRecordingStateEvent) GetData() libs.IEvent {
	return e
}
