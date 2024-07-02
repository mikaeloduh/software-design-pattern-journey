package entity

import (
	"fmt"
	"io"
)

type ChatRoom struct {
	Writer     io.Writer
	Waterball  *Waterball
	TagService func(TagEvent)
	observers  []IChatRoomObserver
}

func (c *ChatRoom) Send(sender IMember, m Message) {
	_, _ = fmt.Fprint(c.Writer, fmt.Sprintf("%s: %s\n", sender.Id(), m.Content))

	if len(m.Tags) != 0 {
		for _, tag := range m.Tags {
			c.Waterball.TagService(TagEvent{
				TaggedBy: sender,
				TaggedTo: tag,
				Message:  m,
			})
		}
	}

	c.Notify(NewMessageEvent{Sender: sender, Message: m})
}

func (c *ChatRoom) Notify(event NewMessageEvent) {
	for _, observer := range c.observers {
		observer.Update(event)
	}
}

func (c *ChatRoom) Register(observer IChatRoomObserver) {
	c.observers = append(c.observers, observer)
}
