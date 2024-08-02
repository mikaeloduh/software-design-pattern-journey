package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type InteractingState struct {
	bot       *Bot
	waterball *service.Waterball
	libs.SuperState
	UnimplementedBotState
	talkCount int
}

func NewInteractingState(waterball *service.Waterball, bot *Bot) *InteractingState {
	return &InteractingState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *InteractingState) GetState() libs.IState {
	return s
}

func (s *InteractingState) OnNewMessage(event service.NewMessageEvent) {
	line := []string{"Hi hi", "I like your idea!"}
	s.waterball.ChatRoom.Send(service.NewMessage(s.bot, line[s.talkCount%len(line)], event.Sender))

	s.talkCount++
}

func (s *InteractingState) OnNewPost(event service.NewPostEvent) {
	s.waterball.Forum.Comment(event.PostId, service.Comment{Member: s.bot, Content: "How do you guys think about it?"})
}
