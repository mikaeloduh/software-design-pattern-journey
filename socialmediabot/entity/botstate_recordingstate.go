package entity

import (
	"fmt"
	"socialmediabot/libs"
)

// RecordingState
type RecordingState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
	record *Record
}

func NewRecordingState(waterball *Waterball, bot *Bot) *RecordingState {
	return &RecordingState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *RecordingState) GetState() libs.IState {
	return s
}

func (s *RecordingState) Enter(_ libs.IEvent) {
	s.record = &Record{}
}

func (s *RecordingState) Exit() {
	s.replayRecord()
}

func (s *RecordingState) OnNewMessage(_ NewMessageEvent) {
}

func (s *RecordingState) OnSpeak(event SpeakEvent) {
	s.record.AddContent(event.Content)
}

func (s *RecordingState) OnBroadcastStop(_ BroadcastStopEvent) {
	s.bot.fsm.Trigger(ExitRecordingStateEvent{})
}

func (s *RecordingState) replayRecord() {
	s.waterball.ChatRoom.Send(NewMessage(s.bot, fmt.Sprintf("[Record Replay] %s", s.record.GetContent())))
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
