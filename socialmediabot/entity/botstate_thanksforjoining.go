package entity

import (
	"fmt"
	"socialmediabot/libs"
	"time"
)

type ThanksForJoiningState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
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
		if num := len(s.bot.Winners); num > 1 {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, "Tie!"))
		} else if num == 1 {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, fmt.Sprintf("The winner is %s", s.bot.Winners[0])))
		} else {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, "Something went wrong"))
		}

		_ = s.waterball.Broadcast.StopBroadcasting(s.bot)
	} else {
		if num := len(s.bot.Winners); num > 1 {
			s.waterball.ChatRoom.Send(NewMessage(s.bot, "Tie!"))
		} else if num == 1 {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, fmt.Sprintf("The winner is %s", s.bot.Winners[0])))
		} else {
			s.waterball.Broadcast.Transmit(NewSpeak(s.bot, "Something went wrong"))
		}
	}

	s.waterball.timer.Sleep(5 * time.Second) // Simulate task processing

	s.bot.fsm.Trigger(ExitThanksForJoiningStateEvent{})
}

func (s *ThanksForJoiningState) Exit() {
	s.bot.Winners = nil
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
