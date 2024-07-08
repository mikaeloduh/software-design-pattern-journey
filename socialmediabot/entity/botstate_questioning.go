package entity

import "socialmediabot/libs"

type QuestioningState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
	talkCount int
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
	s.askQuestion()
}

func (s *QuestioningState) OnNewMessage(event NewMessageEvent) {
	s.validateAnswer(event.Message.Content, event.Sender)
}

/// privates

func (s *QuestioningState) askQuestion() {
	s.waterball.ChatRoom.Send(s.bot, NewMessage(s.getQuestion().question))
}

func (s *QuestioningState) validateAnswer(answer string, sender Taggable) {
	if answer == s.getQuestion().answer {
		s.waterball.ChatRoom.Send(s.bot, NewMessage("Congrats! you got the answer!", sender))
		s.talkCount++

		s.askQuestion()
	}
}

type Question struct {
	question string
	answer   string
}

func (s *QuestioningState) getQuestion() Question {
	questions := []Question{
		{"請問哪個 SQL 語句用於選擇所有的行？\nA) SELECT *\nB) SELECT ALL\nC) SELECT ROWS\nD) SELECT DATA", "A"},
		{"請問哪個 CSS 屬性可用於設置文字的顏色？\nA) text-align\nB) font-size\nC) color\nD) padding", "C"},
		{"請問在計算機科學中，「XML」代表什麼？\nA) Extensible Markup Language\nB) Extensible Modeling Language\nC) Extended Markup Language\nD) Extended Modeling Language", "A"},
	}

	return questions[s.talkCount%len(questions)]
}
