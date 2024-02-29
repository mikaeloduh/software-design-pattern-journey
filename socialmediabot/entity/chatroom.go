package entity

import (
	"fmt"
	"io"
)

type ChatRoom struct {
	Writer io.Writer
}

func (c *ChatRoom) Send(b Member, m Message) {
	_, _ = fmt.Fprint(c.Writer, fmt.Sprintf("%s: %s", b.Id, m.Content))
}
