package entity

import (
	"fmt"
	"io"
)

type ChatRoom struct {
	Writer     io.Writer
	Waterball  *Waterball
	TagService func(TagEvent)
	observers  []INewMessageObserver
}

func (c *ChatRoom) Send(sender IMember, message Message) {
	_, _ = fmt.Fprint(c.Writer, fmt.Sprintf("%s: %s\n", sender.Id(), message.Content))

	for _, tag := range message.Tags {
		c.Waterball.TagOnlineMember(TagEvent{
			TaggedBy: sender,
			TaggedTo: tag,
			Message:  message,
		})
	}

	c.Notify(NewMessageEvent{Sender: sender, Message: message})
}

func (c *ChatRoom) Notify(event NewMessageEvent) {
	for _, observer := range c.observers {
		observer.Update(event)
	}
}

func (c *ChatRoom) Register(observer INewMessageObserver) {
	c.observers = append(c.observers, observer)
}
