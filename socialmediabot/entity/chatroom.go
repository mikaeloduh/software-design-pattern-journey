package entity

import (
	"fmt"
	"io"
)

type ChatRoom struct {
	Writer io.Writer
}

func (c *ChatRoom) Send(b IMember, m Message) {
	_, _ = fmt.Fprint(c.Writer, fmt.Sprintf("%s: %s", b.Id(), m.Content))
	if len(m.Tags) != 0 {
		for _, tag := range m.Tags {
			tag.Update()
		}
	}
}
