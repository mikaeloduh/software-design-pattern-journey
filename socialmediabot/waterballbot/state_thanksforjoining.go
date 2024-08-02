package waterballbot

import (
	"fmt"
	"time"

	"github.com/benbjohnson/clock"

	"socialmediabot/libs"
	"socialmediabot/service"
)

type ThanksForJoiningState struct {
	bot *Bot
	libs.SuperState
	UnimplementedBotState
	timer *clock.Timer
}

func NewThanksForJoiningState(bot *Bot) *ThanksForJoiningState {
	return &ThanksForJoiningState{
		bot:        bot,
		SuperState: libs.SuperState{},
	}
}

func (s *ThanksForJoiningState) Enter(event libs.IEvent) {
	winners := event.(ExitQuestioningStateEvent).Winners

	if err := s.bot.waterball.Broadcast.GoBroadcasting(s.bot); err == nil {
		if num := len(winners); num > 1 || num == 0 {
			s.bot.waterball.Broadcast.Transmit(service.NewSpeak(s.bot, "Tie!"))
		} else if num == 1 {
			s.bot.waterball.Broadcast.Transmit(service.NewSpeak(s.bot, fmt.Sprintf("The winner is %s", winners[0])))
		} else {
			s.bot.waterball.Broadcast.Transmit(service.NewSpeak(s.bot, "Something went wrong"))
		}

		_ = s.bot.waterball.Broadcast.StopBroadcasting(s.bot)
	} else {
		if num := len(winners); num > 1 || num == 0 {
			s.bot.waterball.ChatRoom.Send(service.NewMessage(s.bot, "Tie!"))
		} else if num == 1 {
			s.bot.waterball.Broadcast.Transmit(service.NewSpeak(s.bot, fmt.Sprintf("The winner is %s", winners[0])))
		} else {
			s.bot.waterball.Broadcast.Transmit(service.NewSpeak(s.bot, "Something went wrong"))
		}
	}

	s.timer = s.bot.waterball.Clock.AfterFunc(5*time.Second, func() { s.bot.Update(ExitThanksForJoiningStateEvent{}) })
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

func (s *ThanksForJoiningState) OnSpeak(_ service.SpeakEvent) {
}

func (s *ThanksForJoiningState) OnBroadcastStop(_ service.BroadcastStopEvent) {
}

// ExitThanksForJoiningStateEvent
type ExitThanksForJoiningStateEvent struct {
}

func (e ExitThanksForJoiningStateEvent) GetData() libs.IEvent {
	return e
}
