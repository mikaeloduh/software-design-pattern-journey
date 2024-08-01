package entity

import (
	"fmt"
	"time"

	"github.com/benbjohnson/clock"

	"socialmediabot/libs"
)

type ThanksForJoiningState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
	timer *clock.Timer
}

func NewThanksForJoiningState(waterball *Waterball, bot *Bot) *ThanksForJoiningState {
	return &ThanksForJoiningState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *ThanksForJoiningState) Enter() {
	err := s.waterball.Broadcast.GoBroadcasting(s.bot)
	if err == nil {
		if num := len(s.bot.Winners); num > 1 || num == 0 {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, "Tie!"))
		} else if num == 1 {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, fmt.Sprintf("The winner is %s", s.bot.Winners[0])))
		} else {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, "Something went wrong"))
		}

		_ = s.waterball.Broadcast.StopBroadcasting(s.bot)
	} else {
		if num := len(s.bot.Winners); num > 1 || num == 0 {
			s.waterball.ChatRoom.Send(NewMessage(s.bot, "Tie!"))
		} else if num == 1 {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, fmt.Sprintf("The winner is %s", s.bot.Winners[0])))
		} else {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, "Something went wrong"))
		}
	}

	s.timer = s.waterball.Clock.AfterFunc(5*time.Second, func() { s.bot.fsm.Trigger(ExitThanksForJoiningStateEvent{}) })
}

func (s *ThanksForJoiningState) Exit() {
	s.bot.Winners = nil
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

func (s *ThanksForJoiningState) OnSpeak(event SpeakEvent) {
}

func (s *ThanksForJoiningState) OnBroadcastStop(event BroadcastStopEvent) {
}

// ExitThanksForJoiningStateEvent
type ExitThanksForJoiningStateEvent struct {
}

func (e ExitThanksForJoiningStateEvent) GetData() libs.IEvent {
	return e
}
