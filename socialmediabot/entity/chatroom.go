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
	TagService func(tag Tag)
	observers  []IChatRoomObserver
}

func (c *ChatRoom) Send(mb IMember, mg Message) {
	_, _ = fmt.Fprint(c.Writer, fmt.Sprintf("%s: %s", mb.Id(), mg.Content))
	c.Tagging(mg)
	c.Notify(NewMessageEvent{sender: mb, message: mg})
}

func (c *ChatRoom) Tagging(mg Message) {
	if len(mg.Tags) != 0 {
		for _, tag := range mg.Tags {
			c.TagService(tag)
		}
	}
}

func (c *ChatRoom) Notify(event NewMessageEvent) {
	for _, observer := range c.observers {
		observer.Update(event)
	}
}
