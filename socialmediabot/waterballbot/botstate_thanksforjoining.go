package waterballbot

import (
	"fmt"
	"time"

	"github.com/benbjohnson/clock"

	"socialmediabot/entity"
	"socialmediabot/libs"
)

type ThanksForJoiningState struct {
	bot       *Bot
	waterball *entity.Waterball
	libs.SuperState
	UnimplementedBotState
	timer *clock.Timer
}

func NewThanksForJoiningState(waterball *entity.Waterball, bot *Bot) *ThanksForJoiningState {
	return &ThanksForJoiningState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *ThanksForJoiningState) Enter(event libs.IEvent) {
	winners := event.(ExitQuestioningStateEvent).Winners

	if err := s.waterball.Broadcast.GoBroadcasting(s.bot); err == nil {
		if num := len(winners); num > 1 || num == 0 {
			s.waterball.Broadcast.Transmit(entity.NewSpeak(s.bot, "Tie!"))
		} else if num == 1 {
			s.waterball.Broadcast.Transmit(entity.NewSpeak(s.bot, fmt.Sprintf("The winner is %s", winners[0])))
		} else {
			s.waterball.Broadcast.Transmit(entity.NewSpeak(s.bot, "Something went wrong"))
		}

		_ = s.waterball.Broadcast.StopBroadcasting(s.bot)
	} else {
		if num := len(winners); num > 1 || num == 0 {
			s.waterball.ChatRoom.Send(entity.NewMessage(s.bot, "Tie!"))
		} else if num == 1 {
			s.waterball.Broadcast.Transmit(entity.NewSpeak(s.bot, fmt.Sprintf("The winner is %s", winners[0])))
		} else {
			s.waterball.Broadcast.Transmit(entity.NewSpeak(s.bot, "Something went wrong"))
		}
	}

	s.timer = s.waterball.Clock.AfterFunc(5*time.Second, func() { s.bot.Update(ExitThanksForJoiningStateEvent{}) })
}

func (s *ThanksForJoiningState) Exit() {
	if !s.timer.Stop() {
		select {
		case <-s.timer.C:
			// Drained the channel successfully
		default:
			// The channel was already empty
		}
	}
}

func (s *ThanksForJoiningState) GetState() libs.IState {
	return s
}

func (s *ThanksForJoiningState) OnSpeak(event entity.SpeakEvent) {
}

func (s *ThanksForJoiningState) OnBroadcastStop(event entity.BroadcastStopEvent) {
}

// ExitThanksForJoiningStateEvent
type ExitThanksForJoiningStateEvent struct {
}

func (e ExitThanksForJoiningStateEvent) GetData() libs.IEvent {
	return e
}
