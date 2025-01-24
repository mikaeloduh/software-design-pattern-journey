package waterballbot

import (
	"socialmediabot/libs"
	"socialmediabot/service"
)

type InteractingState struct {
	bot *Bot
	libs.SuperState[IBotState]
	UnimplementedBotOperation
	talkCount int
}

func NewInteractingState(bot *Bot) *InteractingState {
	return &InteractingState{
		bot:        bot,
		SuperState: libs.SuperState[IBotState]{},
	}
}

func (s *InteractingState) GetState() IBotState {
	return s
}

func (s *InteractingState) OnNewMessage(event service.NewMessageEvent) {
	line := []string{"Hi hi", "I like your idea!"}
	s.bot.waterball.ChatRoom.Send(service.NewMessage(s.bot, line[s.talkCount%len(line)], event.Sender))

	s.talkCount++
}

func (s *InteractingState) OnNewPost(event service.NewPostEvent) {
	s.bot.waterball.Forum.Comment(event.PostId, service.Comment{Member: s.bot, Content: "How do you guys think about it?"})
}
