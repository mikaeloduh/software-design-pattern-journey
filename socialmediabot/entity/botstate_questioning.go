package entity

import "socialmediabot/libs"

type QuestioningState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
}

func NewQuestioningState(waterball *Waterball, bot *Bot) *QuestioningState {
	return &QuestioningState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *QuestioningState) GetState() libs.IState {
	return s
}

func (s *QuestioningState) Enter() {
	s.waterball.ChatRoom.Send(s.bot, NewMessage("KnowledgeKing is started!"))
}
