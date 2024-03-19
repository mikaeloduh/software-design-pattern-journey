package entity

import (
	"fmt"
	"io"
)

type IChatRoomObserver interface {
	Update(event NewMessageEvent)
}

type ChatRoom struct {
	Writer     io.Writer
	TagService func(TagEvent)
	observers  []IChatRoomObserver
}

func (c *ChatRoom) Send(sender IMember, m Message) {
	_, _ = fmt.Fprint(c.Writer, fmt.Sprintf("%s: %s", sender.Id(), m.Content))

	if len(m.Tags) != 0 {
		for _, tag := range m.Tags {
			c.TagService(TagEvent{TaggedBy: sender, TaggedTo: tag, Message: m})
		}
	}

	c.Notify(NewMessageEvent{Sender: sender, Message: m})
}

func (c *ChatRoom) Notify(event NewMessageEvent) {
	for _, observer := range c.observers {
		observer.Update(event)
	}
}
