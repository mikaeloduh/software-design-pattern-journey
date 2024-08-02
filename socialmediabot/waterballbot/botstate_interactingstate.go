package waterballbot

import (
	"socialmediabot/entity"
	"socialmediabot/libs"
)

type InteractingState struct {
	bot       *Bot
	waterball *entity.Waterball
	libs.SuperState
	UnimplementedBotState
	talkCount int
}

func NewInteractingState(waterball *entity.Waterball, bot *Bot) *InteractingState {
	return &InteractingState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *InteractingState) GetState() libs.IState {
	return s
}

func (s *InteractingState) OnNewMessage(event entity.NewMessageEvent) {
	line := []string{"Hi hi", "I like your idea!"}
	s.waterball.ChatRoom.Send(entity.NewMessage(s.bot, line[s.talkCount%len(line)], event.Sender))

	s.talkCount++
}

func (s *InteractingState) OnNewPost(event entity.NewPostEvent) {
	s.waterball.Forum.Comment(event.PostId, entity.Comment{Member: s.bot, Content: "How do you guys think about it?"})
}
