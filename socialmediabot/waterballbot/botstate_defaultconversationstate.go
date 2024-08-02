package waterballbot

import (
	"socialmediabot/entity"
	"socialmediabot/libs"
)

// DefaultConversationState
type DefaultConversationState struct {
	bot       *Bot
	waterball *entity.Waterball
	libs.SuperState
	UnimplementedBotState
	talkCount int
}

func NewDefaultConversationState(waterball *entity.Waterball, bot *Bot) *DefaultConversationState {
	return &DefaultConversationState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *DefaultConversationState) GetState() libs.IState {
	return s
}

func (s *DefaultConversationState) OnNewMessage(event entity.NewMessageEvent) {
	line := []string{"good to hear", "thank you", "How are you"}
	s.waterball.ChatRoom.Send(entity.NewMessage(s.bot, line[s.talkCount%len(line)], event.Sender))

	s.talkCount++
}

func (s *DefaultConversationState) OnNewPost(event entity.NewPostEvent) {
	s.waterball.Forum.Comment(event.PostId, entity.Comment{Member: s.bot, Content: "Nice post"})
}
