package entity

import (
	"time"

	"github.com/benbjohnson/clock"

	"socialmediabot/libs"
)

type QuestioningState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
	talkCount  int
	scoreBoard map[string]int
	quitCh     chan bool
	timer      *clock.Timer
}

func NewQuestioningState(waterball *Waterball, bot *Bot) *QuestioningState {
	return &QuestioningState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
		scoreBoard: make(map[string]int),
		quitCh:     make(chan bool, 1),
	}
}

func (s *QuestioningState) GetState() libs.IState {
	return s
}

func (s *QuestioningState) Enter() {
	s.waterball.ChatRoom.Send(NewMessage(s.bot, "KnowledgeKing is started!"))

	s.timer = s.waterball.Clock.AfterFunc(1*time.Hour, s.afterGameEnd)

	s.askQuestion()
}

func (s *QuestioningState) Exit() {
	s.talkCount = 0
	if !s.timer.Stop() {
		select {
		case <-s.timer.C:
			// Drained the channel successfully
		default:
			// The channel was already empty
		}
	}
}

func (s *QuestioningState) OnNewMessage(event NewMessageEvent) {
	s.validateAnswer(event.Message.Content, event.Sender)
}

/// privates

func (s *QuestioningState) askQuestion() {
	s.waterball.ChatRoom.Send(NewMessage(s.bot, s.getQuestion().question))
}

func (s *QuestioningState) validateAnswer(answer string, sender Taggable) {
	if answer == s.getQuestion().answer {
		s.waterball.ChatRoom.Send(NewMessage(s.bot, "Congrats! you got the answer!", sender))
		s.scoreBoard[sender.Id()] += 1
		s.talkCount++

		if s.isGameEnd() {
			s.bot.Winners = findMax(s.scoreBoard)
			s.afterGameEnd()
			return
		}

		s.askQuestion()
	}
}

func (s *QuestioningState) isGameEnd() bool {
	return s.talkCount >= 3
}

func (s *QuestioningState) afterGameEnd() {
	s.bot.fsm.Trigger(ExitQuestioningStateEvent{})
}

// Question
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

func findMax(myMap map[string]int) []string {
	var maxVal int
	var keysWithMaxVal []string
	first := true

	for key, value := range myMap {
		if first || value > maxVal {
			maxVal = value
			keysWithMaxVal = []string{key}
			first = false
		} else if value == maxVal {
			keysWithMaxVal = append(keysWithMaxVal, key)
		}
	}

	return keysWithMaxVal
}

// ExitQuestioningStateEvent
type ExitQuestioningStateEvent struct {
}

func (e ExitQuestioningStateEvent) GetData() libs.IEvent {
	return e
}
