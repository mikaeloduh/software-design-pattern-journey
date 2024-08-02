package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

// DefaultConversationState
type DefaultConversationState struct {
	bot       *Bot
	waterball *service.Waterball
	libs.SuperState
	UnimplementedBotState
	talkCount int
}

func NewDefaultConversationState(waterball *service.Waterball, bot *Bot) *DefaultConversationState {
	return &DefaultConversationState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *DefaultConversationState) GetState() libs.IState {
	return s
}

func (s *DefaultConversationState) OnNewMessage(event service.NewMessageEvent) {
	line := []string{"good to hear", "thank you", "How are you"}
	s.waterball.ChatRoom.Send(service.NewMessage(s.bot, line[s.talkCount%len(line)], event.Sender))

	s.talkCount++
}

func (s *DefaultConversationState) OnNewPost(event service.NewPostEvent) {
	s.waterball.Forum.Comment(event.PostId, service.Comment{Member: s.bot, Content: "Nice post"})
}
