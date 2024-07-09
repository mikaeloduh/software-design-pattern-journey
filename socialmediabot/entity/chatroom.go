package entity

import (
	"fmt"
	"io"
)

type ChatRoom struct {
	writer    io.Writer
	waterball *Waterball
	observers []INewMessageObserver
}

func (c *ChatRoom) Send(message Message) {
	_, _ = fmt.Fprint(c.writer, fmt.Sprintf("%s: %s\f", message.Sender.Id(), message.Content))

	c.Notify(NewMessageEvent{Sender: message.Sender, Message: message})

	for _, tag := range message.Tags {
		c.waterball.TagOnlineMember(TagEvent{
			TaggedBy: message.Sender,
			TaggedTo: tag,
			Message:  message,
		})
	}
}

func (c *ChatRoom) Notify(event NewMessageEvent) {
	for _, observer := range c.observers {
		observer.Update(event)
	}
}

func (c *ChatRoom) Register(observer INewMessageObserver) {
	c.observers = append(c.observers, observer)
}
