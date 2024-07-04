package entity

import "socialmediabot/libs"

type RecordingState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
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

func (s *RecordingState) OnNewMessage(_ NewMessageEvent) {
	// do nothing
}
