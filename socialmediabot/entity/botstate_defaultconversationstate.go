package entity

import "socialmediabot/libs"

// DefaultConversationState
type DefaultConversationState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
	talkCount int
}

func NewDefaultConversationState(waterball *Waterball, bot *Bot) *DefaultConversationState {
	return &DefaultConversationState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *DefaultConversationState) GetState() libs.IState {
	return s
}

func (s *DefaultConversationState) OnNewMessage(event NewMessageEvent) {
	line := []string{"good to hear", "thank you", "How are you"}
	s.waterball.ChatRoom.Send(s.bot, Message{Content: line[s.talkCount%len(line)]})

	s.talkCount++
}

func (s *DefaultConversationState) OnNewPost(event NewPostEvent) {
	s.waterball.Forum.Comment(event.PostId, Comment{Member: s.bot, Content: "Nice post"})
}
