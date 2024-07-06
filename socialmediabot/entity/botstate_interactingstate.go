package entity

import "socialmediabot/libs"

type InteractingState struct {
	bot       *Bot
	waterball *Waterball
	libs.SuperState
	UnimplementedBotState
	talkCount int
}

func NewInteractingState(waterball *Waterball, bot *Bot) *InteractingState {
	return &InteractingState{
		bot:        bot,
		waterball:  waterball,
		SuperState: libs.SuperState{},
	}
}

func (s *InteractingState) GetState() libs.IState {
	return s
}

func (s *InteractingState) OnNewMessage(event NewMessageEvent) {
	line := []string{"Hi hi", "I like your idea!"}
	s.waterball.ChatRoom.Send(s.bot, Message{Content: line[s.talkCount%len(line)]})

	s.talkCount++
}

func (s *InteractingState) OnNewPost(event NewPostEvent) {
	s.waterball.Forum.Comment(event.PostId, Comment{Member: s.bot, Content: "How do you guys think about it?"})
}
