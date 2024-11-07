package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

// DefaultConversationState
type DefaultConversationState struct {
	bot *Bot
	libs.SuperState[IBotState]
	UnimplementedBotState
	talkCount int
}

func NewDefaultConversationState(bot *Bot) *DefaultConversationState {
	return &DefaultConversationState{
		bot:        bot,
		SuperState: libs.SuperState[IBotState]{},
	}
}

func (s *DefaultConversationState) GetState() IBotState {
	return s
}

func (s *DefaultConversationState) OnNewMessage(event service.NewMessageEvent) {
	line := []string{"good to hear", "thank you", "How are you"}
	s.bot.waterball.ChatRoom.Send(service.NewMessage(s.bot, line[s.talkCount%len(line)], event.Sender))

	s.talkCount++
}

func (s *DefaultConversationState) OnNewPost(event service.NewPostEvent) {
	s.bot.waterball.Forum.Comment(event.PostId, service.Comment{Member: s.bot, Content: "Nice post"})
}
