package entity

import "socialmediabot/libs"

type RecordingState struct {
	waterball *Waterball
	libs.SuperState[*Bot]
	UnimplementedBotState
}

func NewRecordingState(waterball *Waterball, bot *Bot) *RecordingState {
	return &RecordingState{
		waterball: waterball,
		SuperState: libs.SuperState[*Bot]{
			Subject: bot,
		},
	}
}

func (s *RecordingState) GetState() libs.IState {
	return s
}

func (s *RecordingState) OnNewMessage(_ NewMessageEvent) {
	// do nothing
}
