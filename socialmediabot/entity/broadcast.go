package entity

import (
	"fmt"
	"io"
	"socialmediabot/libs"
)

type IBroadcastObserver interface {
	Update(event libs.IEvent)
}

// Broadcast
type Broadcast struct {
	writer    io.Writer
	speaker   IMember
	observers []IBroadcastObserver
}

func (c *Broadcast) GoBroadcasting(speaker IMember) error {
	if c.speaker != nil {
		return fmt.Errorf("broacashing is already running")
	}

	c.speaker = speaker
	_, _ = fmt.Fprint(c.writer, fmt.Sprintf("%s go broadcasting...\f", speaker.Id()))

	c.Notify(GoBroadcastingEvent{})

	return nil
}

func (c *Broadcast) StopBroadcasting(speaker IMember) error {
	if speaker != c.speaker {
		return fmt.Errorf("cannot stop broadcasting, not a member")
	}

	c.speaker = nil
	_, _ = fmt.Fprint(c.writer, fmt.Sprintf("%s stop broadcasting\f", speaker.Id()))

	c.Notify(BroadcastStopEvent{})

	return nil
}

func (c *Broadcast) Transmit(speak Speak) {
	if speak.Speaker != c.speaker {
		return
	}

	_, _ = fmt.Fprint(c.writer, fmt.Sprintf("%s speaking: %s\f", speak.Speaker.Id(), speak.Content))

	c.Notify(SpeakEvent{Speaker: speak.Speaker, Content: speak.Content})
}

func (c *Broadcast) Notify(event libs.IEvent) {
	for _, observer := range c.observers {
		observer.Update(event)
	}
}

func (c *Broadcast) Register(observer IBroadcastObserver) {
	c.observers = append(c.observers, observer)
}

// Speak
type Speak struct {
	Speaker IMember
	Content string
}

func NewSpeak(speaker IMember, content string) Speak {
	return Speak{Speaker: speaker, Content: content}
}
