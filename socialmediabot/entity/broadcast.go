package entity

import (
	"fmt"
	"io"
)

// Broadcast
type Broadcast struct {
	writer  io.Writer
	speaker IMember
}

func (c *Broadcast) GoBroadcasting(speaker IMember) {
	c.speaker = speaker
	_, _ = fmt.Fprint(c.writer, fmt.Sprintf("%s go broadcasting...\f", speaker.Id()))
}

func (c *Broadcast) Transmit(speak Speak) {
	if speak.Speaker != c.speaker {
		return
	}

	_, _ = fmt.Fprint(c.writer, fmt.Sprintf("%s speaking: %s\f", speak.Speaker.Id(), speak.Content))
}

// Speak
type Speak struct {
	Speaker IMember
	Content string
}

func NewSpeak(speaker IMember, content string) Speak {
	return Speak{Speaker: speaker, Content: content}
}
