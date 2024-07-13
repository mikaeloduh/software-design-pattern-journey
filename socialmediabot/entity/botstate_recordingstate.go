package entity

import "socialmediabot/libs"

type Record struct {
	Content string
}

func (r *Record) AddContent(content string) {
	if len(r.Content) == 0 {
		r.Content = content
	} else {
		r.Content = r.Content + "\n" + content
	}
}

type RecordingState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
	record Record
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

func (s *RecordingState) Enter() {
	s.record = Record{}
}

func (s *RecordingState) OnNewMessage(_ NewMessageEvent) {
	// do nothing
}

func (s *RecordingState) OnSpeak(event SpeakEvent) {
	s.record.AddContent(event.Content)
}

func (s *RecordingState) OnBroadcastStop(event BroadcastStopEvent) {
	replay := "[Record Replay] " + s.record.Content
	s.waterball.ChatRoom.Send(NewMessage(s.bot, replay))
}
